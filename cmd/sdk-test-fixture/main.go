// sdk-test-fixture is a small helper binary used by the TypeScript SDK
// integration tests.  It connects to a running Anvil + deployer testnet,
// bootstraps the default Anvil node keys, and provides three actions:
//
//	--action=create  (default) creates a Live single-participant DKG epoch
//	                 (finalizeEpoch stores all MaxK pool keys at once) and
//	                 writes JSON to stdout:
//	                   {"epochId":"0x…","share":"<decimal>",
//	                    "shares":["<decimal>",…],
//	                    "poolKey":{"x":"<decimal>","y":"<decimal>"},
//	                    "activatedKeys":1}
//	                 `shares[j]` is the polynomial share value held by
//	                 participant 1 under pool key j, one entry per key the
//	                 caller asked for with `--keys` (every key of a Live
//	                 epoch is claimable; the flag only sizes the output);
//	                 `share` is `shares[0]`. Every key of the pool is
//	                 dealt from its own polynomial, so the test must pass the
//	                 share of the key its application claimed back to the
//	                 decrypt actions for the helper to drive partial
//	                 decryption + combine over an SDK-submitted ciphertext.
//	                 `poolKey` is P_0, the committee key an application
//	                 registered against key 0 encrypts under (plus PK_org
//	                 when it is organizer-locked).
//
//	--action=decrypt drives the threshold-decryption flow for a
//	                 ciphertext that the SDK already submitted on-chain:
//	                 builds the partial decryption proof with its Merkle
//	                 path against the pool key's share root (member 1's
//	                 leaf of the committee-wide tree), calls
//	                 submitPartialDecryption and finally combineDecryption.
//	                 Required additional flags:
//	                   --epoch-id, --aid, --ciphertext-index, --share
//	                 --org-secret is required for an organizer-locked
//	                 application and must be omitted (or 0) for an
//	                 automatic one.
//	                 For an organizer-locked application this action must
//	                 run AFTER revealOrganizerSecret: the contract refuses
//	                 every partial of a sealed application
//	                 (OrganizerSecretNotRevealed), so there is nothing to
//	                 combine before the reveal. The SDK tests reveal first;
//	                 should the application still be sealed, the helper
//	                 publishes --org-secret itself before building the
//	                 partial, and it refuses a secret that differs from an
//	                 already revealed one.
//	                 Outputs `{"ok":true}` on success.
//
//	--action=prepare-combine does the same but stops before the combine and
//	                 emits the calldata, so the SDK writer issues the
//	                 combineDecryption transaction itself. Same reveal
//	                 precondition as `decrypt`.
//
// The TypeScript tests use these together to verify the full epoch-trip
// (encrypt -> submitCiphertext -> partial decrypt -> combine -> getPlaintext)
// without having to generate Groth16 proofs in TypeScript.
//
// Flags (all can be provided as CLI flags or their env-var equivalents):
//
//	--rpc-url        RPC_URL env var or DAVINCI_DKG_TEST_RPC_URL
//	--addresses-file path to addresses.env served by the deployer container
package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	flag "github.com/spf13/pflag"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/log"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
)

type point struct {
	X string `json:"x"`
	Y string `json:"y"`
}

type fixtureResult struct {
	EpochID string `json:"epochId"`
	// Share is participant 1's share of pool key 0, i.e. Shares[0].
	Share string `json:"share"`
	// Shares[j] is participant 1's share of pool key j, one per key the
	// caller asked for; every key of a Live epoch is claimable.
	Shares  []string `json:"shares"`
	PoolKey point    `json:"poolKey"`
	// ActivatedKeys is len(Shares): the keys the output describes. The name
	// predates batched finalization, when keys were proven one at a time.
	ActivatedKeys int `json:"activatedKeys"`
}

type decryptResult struct {
	OK bool `json:"ok"`
}

// prepareCombineResult is emitted by `--action=prepare-combine` so the SDK
// can drive the on-chain combineDecryption call itself. All bytes are
// 0x-prefixed lower-case hex.
type prepareCombineResult struct {
	CombineHash string `json:"combineHash"`
	Plaintext   string `json:"plaintext"`  // decimal
	Transcript  string `json:"transcript"` // 0x-hex
	Proof       string `json:"proof"`      // 0x-hex
	Input       string `json:"input"`      // 0x-hex
}

func main() {
	var rpcURL string
	var addressesFile string
	var action string
	var roundIDHex string
	var aidHex string
	var ciphertextIndex int
	var shareDec string
	var orgSecretHex string
	var keys int

	flag.StringVar(&rpcURL, "rpc-url", os.Getenv("DAVINCI_DKG_TEST_RPC_URL"),
		"RPC URL of the Anvil testnet")
	flag.StringVar(&addressesFile, "addresses-file", os.Getenv("DAVINCI_DKG_TEST_ADDRESSES"),
		"path to addresses.env file (as served by the deployer container)")
	flag.StringVar(&action, "action", "create",
		"action to perform: 'create' (default), 'decrypt' or 'prepare-combine'")
	flag.IntVar(&keys, "keys", 1,
		"(create) how many of the epoch's pool keys to emit shares for (1..MaxK; every key is claimable once Live)")
	flag.StringVar(&roundIDHex, "epoch-id", "",
		"(decrypt) epoch id as a 0x-prefixed 12-byte hex string")
	flag.IntVar(&ciphertextIndex, "ciphertext-index", 0,
		"(decrypt) ciphertext index to combine (must be > 0)")
	flag.StringVar(&aidHex, "aid", "",
		"(decrypt) application id as a 0x-prefixed 32-byte hex string")
	flag.StringVar(&shareDec, "share", "",
		"(decrypt) participant 1's share of the pool key the application claimed "+
			"(`shares[poolIndex]` of the create output), decimal")
	flag.StringVar(&orgSecretHex, "org-secret", "",
		"(decrypt) organizer secret of an organizer-locked application, revealed on chain before any partial "+
			"(the helper reveals it when the application is still sealed); omit for automatic ones")
	flag.Parse()

	if rpcURL == "" {
		fmt.Fprintln(os.Stderr, "error: --rpc-url is required")
		os.Exit(1)
	}
	if addressesFile == "" {
		fmt.Fprintln(os.Stderr, "error: --addresses-file is required")
		os.Exit(1)
	}

	log.Init("info", "stderr", nil) // keep stdout clean for JSON output

	addressesContent, err := os.ReadFile(addressesFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read addresses file: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	services, cleanup, err := helpers.NewTestServicesFromExternal(ctx, rpcURL, addressesContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: setup services: %v\n", err)
		os.Exit(1)
	}
	defer cleanup()

	switch action {
	case "create":
		if keys < 1 || keys > ccommon.MaxK {
			fmt.Fprintf(os.Stderr, "error: --keys must be in [1, %d]\n", ccommon.MaxK)
			os.Exit(1)
		}
		result, err := helpers.CreateSDKTestFixture(ctx, services)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: create fixture: %v\n", err)
			os.Exit(1)
		}
		poolKey := result.PoolKey(0)
		shares := make([]string, keys)
		for keyIndex := range shares {
			share, err := result.ParticipantShare(uint8(keyIndex), 1)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			shares[keyIndex] = share.String()
		}
		emit(fixtureResult{
			EpochID:       fmt.Sprintf("0x%x", result.EpochID),
			Share:         shares[0],
			Shares:        shares,
			PoolKey:       point{X: poolKey.X.String(), Y: poolKey.Y.String()},
			ActivatedKeys: len(shares),
		})

	case "decrypt":
		epochID, aid, share, skOrg := decryptArgs(roundIDHex, aidHex, shareDec, orgSecretHex, ciphertextIndex)
		if err := helpers.CombineSingleParticipantDecryption(
			ctx, services, epochID, aid, uint16(ciphertextIndex), share, skOrg,
		); err != nil {
			fmt.Fprintf(os.Stderr, "error: combine decryption: %v\n", err)
			os.Exit(1)
		}
		emit(decryptResult{OK: true})

	case "prepare-combine":
		epochID, aid, share, skOrg := decryptArgs(roundIDHex, aidHex, shareDec, orgSecretHex, ciphertextIndex)
		payload, err := helpers.PrepareSingleParticipantCombinePayload(
			ctx, services, epochID, aid, uint16(ciphertextIndex), share, skOrg,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: prepare combine: %v\n", err)
			os.Exit(1)
		}
		emit(prepareCombineResult{
			CombineHash: "0x" + hex.EncodeToString(payload.CombineHash[:]),
			Plaintext:   payload.Plaintext.String(),
			Transcript:  "0x" + hex.EncodeToString(payload.Transcript),
			Proof:       "0x" + hex.EncodeToString(payload.Proof),
			Input:       "0x" + hex.EncodeToString(payload.Input),
		})

	default:
		fmt.Fprintf(os.Stderr, "error: unknown --action %q (create, decrypt or prepare-combine)\n", action)
		os.Exit(1)
	}
}

// emit writes the JSON result to stdout, or exits.
func emit(value any) {
	encoded, err := json.Marshal(value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: marshal result: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(encoded))
}

// decryptArgs validates and parses the flags both decrypt actions share.
func decryptArgs(
	roundIDHex, aidHex, shareDec, orgSecretHex string, ciphertextIndex int,
) ([12]byte, [32]byte, *big.Int, *big.Int) {
	if roundIDHex == "" {
		fmt.Fprintln(os.Stderr, "error: --epoch-id is required")
		os.Exit(1)
	}
	if ciphertextIndex <= 0 || ciphertextIndex > 0xffff {
		fmt.Fprintln(os.Stderr, "error: --ciphertext-index must be in (0, 65535]")
		os.Exit(1)
	}
	if shareDec == "" {
		fmt.Fprintln(os.Stderr, "error: --share is required")
		os.Exit(1)
	}
	share, ok := new(big.Int).SetString(shareDec, 10)
	if !ok {
		fmt.Fprintf(os.Stderr, "error: --share %q is not a valid decimal\n", shareDec)
		os.Exit(1)
	}
	return mustEpochID(roundIDHex), mustAid(aidHex), share, organizerSecret(orgSecretHex)
}

// mustEpochID parses a 0x-prefixed 12-byte epoch id or exits.
func mustEpochID(value string) [12]byte {
	raw, err := hex.DecodeString(strings.TrimPrefix(value, "0x"))
	if err != nil || len(raw) != 12 {
		fmt.Fprintf(os.Stderr, "error: --epoch-id must be 0x-prefixed 12-byte hex, got %q\n", value)
		os.Exit(1)
	}
	var epochID [12]byte
	copy(epochID[:], raw)
	return epochID
}

// mustAid parses a 0x-prefixed application id (right-aligned into 32 bytes)
// or exits. Every ciphertext belongs to a registered application, so this is
// required for the decrypt actions.
func mustAid(value string) [32]byte {
	raw, err := hex.DecodeString(strings.TrimPrefix(value, "0x"))
	if err != nil || len(raw) == 0 || len(raw) > 32 {
		fmt.Fprintf(os.Stderr, "error: --aid must be 1..32 bytes of hex, got %q\n", value)
		os.Exit(1)
	}
	var aid [32]byte
	copy(aid[32-len(raw):], raw)
	if aid == ([32]byte{}) {
		fmt.Fprintln(os.Stderr, "error: --aid must be non-zero")
		os.Exit(1)
	}
	return aid
}

// organizerSecret parses the revealed organizer scalar, or returns 0 for an
// automatic application (which has no organizer half at all).
func organizerSecret(value string) *big.Int {
	trimmed := strings.TrimPrefix(value, "0x")
	if trimmed == "" {
		return big.NewInt(0)
	}
	sk, ok := new(big.Int).SetString(trimmed, 16)
	if !ok || sk.Sign() < 0 {
		fmt.Fprintf(os.Stderr, "error: --org-secret must be a hex scalar, got %q\n", value)
		os.Exit(1)
	}
	return sk
}
