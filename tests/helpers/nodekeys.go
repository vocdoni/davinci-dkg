package helpers

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	dkghash "github.com/vocdoni/davinci-dkg/crypto/hash"
)

// localNodeKeyDerivationDomain must match bjjKeyDomain in
// cmd/davinci-dkg-node/node.go so that keys derived here match the node.
const localNodeKeyDerivationDomain = "davinci-dkg/bjj-key/v1"

func deterministicNodeKeyMaterial(privateKey string) (*big.Int, *big.Int, *big.Int, error) {
	preimage := append(common.FromHex(privateKey), []byte(localNodeKeyDerivationDomain)...)
	digest := ethcrypto.Keccak256(preimage)
	lo := new(big.Int).SetBytes(digest[:16])
	hi := new(big.Int).SetBytes(digest[16:])
	secret, err := dkghash.HashFieldElements(lo, hi)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("poseidon hash: %w", err)
	}
	secret.Mod(secret, group.ScalarField())
	if secret.Sign() == 0 {
		secret.SetInt64(1)
	}

	publicKey := group.NewPoint()
	publicKey.ScalarBaseMult(secret)
	encoded := group.Encode(publicKey)
	if encoded.X == nil || encoded.Y == nil {
		return nil, nil, nil, fmt.Errorf("encode deterministic node key")
	}
	return encoded.X, encoded.Y, secret, nil
}

func bootstrapLocalNodeKeys(ctx context.Context, services *TestServices) error {
	for _, privateKey := range DefaultAnvilPrivateKeys {
		actor, err := services.ActorFromPrivateKey(privateKey)
		if err != nil {
			return fmt.Errorf("actor from private key: %w", err)
		}

		expectedX, expectedY, _, err := deterministicNodeKeyMaterial(privateKey)
		if err != nil {
			return fmt.Errorf("derive deterministic node key for %s: %w", actor.Address().Hex(), err)
		}

		node, err := services.Contracts.GetNode(ctx, actor.Address())
		if err != nil {
			return fmt.Errorf("get node for %s: %w", actor.Address().Hex(), err)
		}
		if node.Status != 0 &&
			node.Operator == actor.Address() &&
			node.PubX != nil && node.PubY != nil &&
			node.PubX.Cmp(expectedX) == 0 &&
			node.PubY.Cmp(expectedY) == 0 {
			continue
		}

		auth, err := actor.TxManager.NewTransactOpts(ctx)
		if err != nil {
			return fmt.Errorf("tx opts for %s: %w", actor.Address().Hex(), err)
		}

		var txHash common.Hash
		if node.Status == 0 {
			tx, err := actor.Registry.RegisterKey(auth, expectedX, expectedY)
			if err != nil {
				return fmt.Errorf("register key for %s: %w", actor.Address().Hex(), err)
			}
			txHash = tx.Hash()
		} else {
			tx, err := actor.Registry.UpdateKey(auth, expectedX, expectedY)
			if err != nil {
				return fmt.Errorf("update key for %s: %w", actor.Address().Hex(), err)
			}
			txHash = tx.Hash()
		}

		if err := actor.TxManager.WaitTxByHash(txHash, DefaultTxTimeout); err != nil {
			return fmt.Errorf("wait key registration for %s: %w", actor.Address().Hex(), err)
		}
	}
	return nil
}
