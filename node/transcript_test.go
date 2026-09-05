package node

import (
	"bytes"
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	nodetypes "github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

// The node recovers other members' encrypted shares from submitContribution
// calldata, so the decoder must track the contract's real ABI layout and the
// compact transcript's (t, n)-dependent offsets.
func TestDecodeContributionFollowsSubmitContributionABI(t *testing.T) {
	c := qt.New(t)
	layout, err := contribution.NewLayout(2, 3)
	c.Assert(err, qt.IsNil)
	words := make([]*big.Int, layout.Words())
	for i := range words {
		words[i] = big.NewInt(int64(1000 + i))
	}
	for i := range layout.CommitteeSize {
		words[layout.RecipientIndexOffset(i)] = big.NewInt(int64(i + 1))
	}
	words[layout.EphemeralOffset(2)], words[layout.EphemeralOffset(2)+1] = big.NewInt(11), big.NewInt(13)
	// The masked share the last pool key sends member 3.
	words[layout.MaskedShareOffset(contribution.MaxKeys-1, 2)] = big.NewInt(17)
	transcript, err := encodeWords(words...)
	c.Assert(err, qt.IsNil)

	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	c.Assert(err, qt.IsNil)
	calldata, err := parsed.Pack("submitContribution",
		[12]byte{1}, uint16(1), [32]byte{2}, [32]byte{3}, transcript, []byte{0xaa}, []byte{0xbb})
	c.Assert(err, qt.IsNil)

	tr, err := decodeContribution(calldata, layout)
	c.Assert(err, qt.IsNil)
	c.Assert(tr.RecipientIndexes[2].Int64(), qt.Equals, int64(3))
	c.Assert(tr.Ephemerals[2].X.Int64(), qt.Equals, int64(11))
	c.Assert(tr.Ephemerals[2].Y.Int64(), qt.Equals, int64(13))
	c.Assert(tr.MaskedShares[contribution.MaxKeys-1][2].Int64(), qt.Equals, int64(17))
	c.Assert(tr.Commitments[0], qt.HasLen, 2)

	// A cache miss (nil calldata) and a transcript of another policy's size
	// are errors, never panics, so the caller falls back to the chain.
	_, err = decodeContribution(nil, layout)
	c.Assert(err, qt.IsNotNil)
	other, err := contribution.NewLayout(2, 4)
	c.Assert(err, qt.IsNil)
	_, err = decodeContribution(calldata, other)
	c.Assert(err, qt.IsNotNil)
}

// fakeContribChain serves one epoch's contribution records over eth_call and
// the ContributionSubmitted log / transaction the calldata scan reads, so
// buildPrivateShare runs against it without an RPC endpoint.
type fakeContribChain struct {
	fakeLogChain
	head    uint64
	txs     map[common.Hash]*ethtypes.Transaction
	records map[common.Address]gtypes.DKGTypesContributionRecord
}

func (f *fakeContribChain) BlockNumber(context.Context) (uint64, error) { return f.head, nil }

func (f *fakeContribChain) TransactionByHash(_ context.Context, h common.Hash) (*ethtypes.Transaction, bool, error) {
	tx, ok := f.txs[h]
	if !ok {
		return nil, false, ethereum.NotFound
	}
	return tx, false, nil
}

func (f *fakeContribChain) CallContract(_ context.Context, call ethereum.CallMsg, _ *big.Int) ([]byte, error) {
	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	if len(call.Data) < 4 {
		return nil, errFakeUnsupported
	}
	method, err := parsed.MethodById(call.Data[:4])
	if err != nil || method.Name != "getContribution" {
		return nil, errFakeUnsupported
	}
	args, err := method.Inputs.Unpack(call.Data[4:])
	if err != nil {
		return nil, err
	}
	contributor, _ := args[1].(common.Address)
	return method.Outputs.Pack(f.records[contributor])
}

// contributionLog builds a ContributionSubmitted log as the contract emits
// it, so the abigen iterator can unpack it.
func contributionLog(t *testing.T, epochID [12]byte, contributor common.Address, block uint64, tx *ethtypes.Transaction) ethtypes.Log {
	t.Helper()
	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	qt.Assert(t, err, qt.IsNil)
	ev := parsed.Events["ContributionSubmitted"]
	topics, err := abi.MakeTopics([]any{ev.ID}, []any{epochID}, []any{contributor})
	qt.Assert(t, err, qt.IsNil)
	data, err := ev.Inputs.NonIndexed().Pack(uint16(1), [32]byte{}, [32]byte{})
	qt.Assert(t, err, qt.IsNil)
	return ethtypes.Log{
		Topics:      []common.Hash{topics[0][0], topics[1][0], topics[2][0]},
		Data:        data,
		BlockNumber: block,
		TxHash:      tx.Hash(),
	}
}

// dealtContribution builds member dealerIdx's genuine compact contribution to
// a committee whose member myIdx holds `secret` (the other members' keys are
// throwaway), exactly as doContribution does, and returns its public inputs,
// transcript words and per-key polynomials.
func dealtContribution(
	t *testing.T, epochID [12]byte, layout contribution.Layout, dealerIdx, myIdx uint16, secret *big.Int,
) (*contribution.PublicInputs, []*big.Int, [][]*big.Int) {
	t.Helper()
	coeffs := make([][]*big.Int, contribution.MaxKeys)
	for j := range coeffs {
		coeffs[j] = make([]*big.Int, layout.Threshold)
		for m := range coeffs[j] {
			coeffs[j][m] = big.NewInt(int64(1000*(j+1) + m + 7))
		}
	}
	idxs := make([]uint16, layout.CommitteeSize)
	keys := make([]nodetypes.NodeKey, layout.CommitteeSize)
	nonces := make([]*big.Int, layout.CommitteeSize)
	for i := range idxs {
		idxs[i] = uint16(i + 1)
		sk := big.NewInt(int64(100 + i))
		if idxs[i] == myIdx {
			sk = secret
		}
		pub := group.NewPoint()
		pub.ScalarBaseMult(sk)
		pt := group.Encode(pub)
		keys[i] = nodetypes.NodeKey{Operator: common.BigToAddress(big.NewInt(int64(i + 1))), PubX: pt.X, PubY: pt.Y}
		nonces[i] = big.NewInt(int64(500 + i))
	}
	_, pi, err := contribution.BuildWitness(contribution.Assignment{
		RoundHash:        roundScalar(epochID),
		Threshold:        uint16(layout.Threshold),
		CommitteeSize:    uint16(layout.CommitteeSize),
		ContributorIndex: dealerIdx,
		Coefficients:     coeffs,
		RecipientIndexes: idxs,
		RecipientKeys:    keys,
		EncryptionNonces: nonces,
	})
	qt.Assert(t, err, qt.IsNil)
	words, err := pi.TranscriptScalars()
	qt.Assert(t, err, qt.IsNil)
	return pi, words, coeffs
}

// packContribution packs submitContribution calldata carrying `words` as the
// transcript under the dealer's genuine public inputs.
func packContribution(t *testing.T, epochID [12]byte, dealerIdx uint16, pi *contribution.PublicInputs, words []*big.Int) []byte {
	t.Helper()
	transcript, err := encodeWords(words...)
	qt.Assert(t, err, qt.IsNil)
	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	qt.Assert(t, err, qt.IsNil)
	calldata, err := parsed.Pack("submitContribution", epochID, dealerIdx,
		common.BigToHash(pi.CommitmentHash), common.BigToHash(pi.ShareHash), transcript, []byte{0xaa}, []byte{0xbb})
	qt.Assert(t, err, qt.IsNil)
	return calldata
}

// A dealer's calldata comes from an RPC body located through an event log,
// so before it becomes part of d_{j,i} (and is cached for good) the node must
// check it against the chain: the decoded commitments have to reproduce the
// dealer's stored commitmentsHash and the recovered share has to open the
// dealer's key-j commitments. One flipped masked-share word is refused, not
// cached and leaves no private share behind; a corrupted cache entry is
// refetched rather than trusted.
func TestBuildPrivateShareRefusesAndDoesNotCacheACorruptedContribution(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	layout, err := contribution.NewLayout(2, 3)
	c.Assert(err, qt.IsNil)
	epochID := [12]byte{7}
	const dealerIdx, myIdx = uint16(1), uint16(2)
	secret := big.NewInt(424242)
	selected := []common.Address{
		common.BigToAddress(big.NewInt(1)), common.BigToAddress(big.NewInt(2)), common.BigToAddress(big.NewInt(3)),
	}
	dealer := selected[dealerIdx-1]
	epoch := epochView{Policy: web3.EpochPolicy{Threshold: 2, CommitteeSize: 3}, SeedBlock: 100}
	slot := poolSlot{epoch: epochID, key: 0}

	pi, words, coeffs := dealtContribution(t, epochID, layout, dealerIdx, myIdx, secret)
	genuine := packContribution(t, epochID, dealerIdx, pi, words)
	want, err := ccommon.EvaluatePolynomialNative(coeffs[0], big.NewInt(int64(myIdx)))
	c.Assert(err, qt.IsNil)

	// The masked share key 0 sends member myIdx, with one bit flipped.
	flipped := make([]*big.Int, len(words))
	copy(flipped, words)
	q := layout.MaskedShareOffset(0, int(myIdx)-1)
	flipped[q] = new(big.Int).Xor(words[q], big.NewInt(1))
	badShare := packContribution(t, epochID, dealerIdx, pi, flipped)

	// Two commitments of key 0 swapped: valid points, another commitmentsHash.
	swapped := make([]*big.Int, len(words))
	copy(swapped, words)
	a, b := layout.CommitmentOffset(0, 0), layout.CommitmentOffset(0, 1)
	swapped[a], swapped[a+1], swapped[b], swapped[b+1] = words[b], words[b+1], words[a], words[a+1]
	badCommitments := packContribution(t, epochID, dealerIdx, pi, swapped)

	chainServing := func(calldata []byte) *fakeContribChain {
		tx := ethtypes.NewTx(&ethtypes.LegacyTx{Data: calldata})
		chain := &fakeContribChain{
			head: 200,
			txs:  map[common.Hash]*ethtypes.Transaction{tx.Hash(): tx},
			records: map[common.Address]gtypes.DKGTypesContributionRecord{dealer: {
				Contributor: dealer, ContributorIndex: dealerIdx, CommitmentsHash: common.BigToHash(pi.CommitmentHash), Accepted: true,
			}},
		}
		chain.logs = []ethtypes.Log{contributionLog(t, epochID, dealer, 150, tx)}
		return chain
	}
	nodeOn := func(chain *fakeContribChain) (*Node, string) {
		m, err := gtypes.NewDKGManager(common.Address{}, chain)
		c.Assert(err, qt.IsNil)
		dir := filepath.Join(t.TempDir(), "contributions")
		return &Node{
			address: selected[myIdx-1], bjjSecret: secret, manager: m,
			privateShares: map[poolSlot]*big.Int{}, ownContribs: map[[12]byte]*savedContrib{},
			contribCache: &contributionCache{dir: dir},
		}, dir
	}

	// The genuine calldata yields f_{1,0}(myIdx), is cached and becomes d_{0,myIdx}.
	chain := chainServing(genuine)
	n, _ := nodeOn(chain)
	share, err := n.buildPrivateShare(ctx, chain, epochID, 0, myIdx, selected, epoch)
	c.Assert(err, qt.IsNil)
	c.Assert(share.Cmp(want), qt.Equals, 0)
	c.Assert(n.privateShares[slot], qt.Not(qt.IsNil))
	cached, ok := n.contribCache.Get(epochID, dealer)
	c.Assert(ok, qt.IsTrue)
	c.Assert(bytes.Equal(cached, genuine), qt.IsTrue)

	// One flipped masked-share word: the share does not open the dealer's
	// commitments, so it is an error, nothing is cached, no share is kept.
	for name, bad := range map[string][]byte{"masked share": badShare, "commitments": badCommitments} {
		chain = chainServing(bad)
		n, dir := nodeOn(chain)
		_, err = n.buildPrivateShare(ctx, chain, epochID, 0, myIdx, selected, epoch)
		c.Assert(err, qt.IsNotNil, qt.Commentf("%s", name))
		if name == "masked share" {
			c.Assert(err, qt.ErrorMatches, ".*share does not match commitments.*")
		} else {
			c.Assert(err, qt.ErrorMatches, ".*commitments hash .* does not match the stored.*")
		}
		_, ok = n.contribCache.Get(epochID, dealer)
		c.Assert(ok, qt.IsFalse, qt.Commentf("%s: corrupted calldata must not be cached", name))
		_, statErr := os.Stat(dir)
		c.Assert(os.IsNotExist(statErr), qt.IsTrue, qt.Commentf("%s: no cache file may be written", name))
		_, ok = n.privateShares[slot]
		c.Assert(ok, qt.IsFalse, qt.Commentf("%s: no private share may be kept", name))
	}

	// A cache entry that fails the same checks is replaced by a fresh fetch.
	chain = chainServing(genuine)
	n, _ = nodeOn(chain)
	n.contribCache.Put(epochID, dealer, badShare)
	share, err = n.buildPrivateShare(ctx, chain, epochID, 0, myIdx, selected, epoch)
	c.Assert(err, qt.IsNil)
	c.Assert(share.Cmp(want), qt.Equals, 0)
	cached, ok = n.contribCache.Get(epochID, dealer)
	c.Assert(ok, qt.IsTrue)
	c.Assert(bytes.Equal(cached, genuine), qt.IsTrue)
}
