package web3

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/types"
)

type testRPCRequest struct {
	ID     json.RawMessage   `json:"id"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}

type testRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Result  any             `json:"result,omitempty"`
	Error   *testRPCError   `json:"error,omitempty"`
}

type testRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func TestNewRejectsMissingAddresses(t *testing.T) {
	c := qt.New(t)

	_, err := New([]string{"http://127.0.0.1:8545"}, types.ContractAddresses{})

	c.Assert(err, qt.Not(qt.IsNil))
	c.Assert(err.Error(), qt.Contains, "manager address is required")
}

func TestVerifierKeyHashes(t *testing.T) {
	c := qt.New(t)

	server := testRPCServer()
	defer server.Close()

	contracts, err := New([]string{server.URL}, types.ContractAddresses{
		Registry:               common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Manager:                common.HexToAddress("0x2000000000000000000000000000000000000002"),
		ContributionVerifier:   common.HexToAddress("0x3000000000000000000000000000000000000003"),
		FinalizeVerifier:       common.HexToAddress("0x4000000000000000000000000000000000000004"),
		PartialDecryptVerifier: common.HexToAddress("0x5000000000000000000000000000000000000005"),
		DecryptCombineVerifier: common.HexToAddress("0x6000000000000000000000000000000000000006"),
		RevealSubmitVerifier:   common.HexToAddress("0x7000000000000000000000000000000000000007"),
		RevealShareVerifier:    common.HexToAddress("0x8000000000000000000000000000000000000008"),
	})
	c.Assert(err, qt.IsNil)

	contributionHash, err := contracts.GetContributionVerifierVKeyHash(context.Background())
	c.Assert(err, qt.IsNil)
	c.Assert(contributionHash, qt.Equals, common.HexToHash("0x1234"))

	partialHash, err := contracts.GetPartialDecryptVerifierVKeyHash(context.Background())
	c.Assert(err, qt.IsNil)
	c.Assert(partialHash, qt.Equals, common.HexToHash("0x5678"))

	finalizeHash, err := contracts.GetFinalizeVerifierVKeyHash(context.Background())
	c.Assert(err, qt.IsNil)
	c.Assert(finalizeHash, qt.Equals, common.HexToHash("0x9abc"))

	decryptCombineHash, err := contracts.GetDecryptCombineVerifierVKeyHash(context.Background())
	c.Assert(err, qt.IsNil)
	c.Assert(decryptCombineHash, qt.Equals, common.HexToHash("0xdef0"))

	revealShareHash, err := contracts.GetRevealShareVerifierVKeyHash(context.Background())
	c.Assert(err, qt.IsNil)
	c.Assert(revealShareHash, qt.Equals, common.HexToHash("0x2468"))
}

func TestGetNodeAndRoundViews(t *testing.T) {
	c := qt.New(t)

	server := testRPCServer()
	defer server.Close()

	contracts, err := New([]string{server.URL}, types.ContractAddresses{
		Registry:               common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Manager:                common.HexToAddress("0x2000000000000000000000000000000000000002"),
		ContributionVerifier:   common.HexToAddress("0x3000000000000000000000000000000000000003"),
		FinalizeVerifier:       common.HexToAddress("0x4000000000000000000000000000000000000004"),
		PartialDecryptVerifier: common.HexToAddress("0x5000000000000000000000000000000000000005"),
		DecryptCombineVerifier: common.HexToAddress("0x6000000000000000000000000000000000000006"),
		RevealSubmitVerifier:   common.HexToAddress("0x7000000000000000000000000000000000000007"),
		RevealShareVerifier:    common.HexToAddress("0x8000000000000000000000000000000000000008"),
	})
	c.Assert(err, qt.IsNil)

	node, err := contracts.GetNode(context.Background(), common.HexToAddress("0x4000000000000000000000000000000000000004"))
	c.Assert(err, qt.IsNil)
	c.Assert(node.Operator, qt.Equals, common.HexToAddress("0x4000000000000000000000000000000000000004"))
	c.Assert(node.PubX.Cmp(big.NewInt(11)), qt.Equals, 0)
	c.Assert(node.Status, qt.Equals, uint8(1))

	var roundID [12]byte
	copy(roundID[:], []byte("round-1"))

	round, err := contracts.GetRound(context.Background(), roundID)
	c.Assert(err, qt.IsNil)
	c.Assert(round.Organizer, qt.Equals, common.HexToAddress("0x5000000000000000000000000000000000000005"))
	c.Assert(round.Policy.Threshold, qt.Equals, uint16(2))
	c.Assert(round.Status, qt.Equals, uint8(3))
	c.Assert(round.ContributionCount, qt.Equals, uint16(2))
	c.Assert(round.RevealedShareCount, qt.Equals, uint16(1))
	// The five hash fields are no longer in storage — consumers read them from
	// RoundFinalized / SecretReconstructed events instead.

	selected, err := contracts.SelectedParticipants(context.Background(), roundID)
	c.Assert(err, qt.IsNil)
	c.Assert(selected, qt.DeepEquals, []common.Address{
		common.HexToAddress("0x5000000000000000000000000000000000000005"),
		common.HexToAddress("0x6000000000000000000000000000000000000006"),
	})

	combined, err := contracts.GetCombinedDecryption(context.Background(), roundID, 1)
	c.Assert(err, qt.IsNil)
	c.Assert(combined.CiphertextIndex, qt.Equals, uint16(1))
	c.Assert(combined.Completed, qt.IsTrue)
	c.Assert(combined.Plaintext.Cmp(big.NewInt(39321)), qt.Equals, 0) // 0x9999 = 39321

	revealed, err := contracts.GetRevealedShare(context.Background(), roundID, common.HexToAddress("0x5000000000000000000000000000000000000005"))
	c.Assert(err, qt.IsNil)
	c.Assert(revealed.Participant, qt.Equals, common.HexToAddress("0x5000000000000000000000000000000000000005"))
	c.Assert(revealed.ParticipantIndex, qt.Equals, uint16(1))
	c.Assert(revealed.ShareHash, qt.Equals, common.HexToHash("0xaaaa"))
	c.Assert(revealed.Accepted, qt.IsTrue)
}

func testRPCServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() { _ = r.Body.Close() }()

		var req testRPCRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := testRPCResponse{
			JSONRPC: "2.0",
			ID:      req.ID,
		}

		switch req.Method {
		case "eth_chainId":
			resp.Result = "0x7a69"
		case "eth_call":
			if len(req.Params) == 0 {
				resp.Error = &testRPCError{Code: -32602, Message: "missing params"}
				break
			}
			var call struct {
				Data  string `json:"data"`
				Input string `json:"input"`
			}
			if err := json.Unmarshal(req.Params[0], &call); err != nil {
				resp.Error = &testRPCError{Code: -32602, Message: err.Error()}
				break
			}
			callData := call.Data
			if callData == "" {
				callData = call.Input
			}
			contributionSelector := "0x" + hex.EncodeToString(managerABI.Methods["getContributionVerifierVKeyHash"].ID)
			finalizeSelector := "0x" + hex.EncodeToString(managerABI.Methods["getFinalizeVerifierVKeyHash"].ID)
			partialSelector := "0x" + hex.EncodeToString(managerABI.Methods["getPartialDecryptVerifierVKeyHash"].ID)
			decryptCombineSelector := "0x" + hex.EncodeToString(managerABI.Methods["getDecryptCombineVerifierVKeyHash"].ID)
			revealShareSelector := "0x" + hex.EncodeToString(managerABI.Methods["getRevealShareVerifierVKeyHash"].ID)
			getNodeSelector := "0x" + hex.EncodeToString(registryABI.Methods["getNode"].ID)
			getRoundSelector := "0x" + hex.EncodeToString(managerABI.Methods["getRound"].ID)
			selectedParticipantsSelector := "0x" + hex.EncodeToString(managerABI.Methods["selectedParticipants"].ID)
			getCombinedDecryptionSelector := "0x" + hex.EncodeToString(managerABI.Methods["getCombinedDecryption"].ID)
			getRevealedShareSelector := "0x" + hex.EncodeToString(managerABI.Methods["getRevealedShare"].ID)
			switch {
			case strings.HasPrefix(callData, contributionSelector):
				resp.Result = "0x" + strings.Repeat("0", 60) + "1234"
			case strings.HasPrefix(callData, finalizeSelector):
				resp.Result = "0x" + strings.Repeat("0", 60) + "9abc"
			case strings.HasPrefix(callData, partialSelector):
				resp.Result = "0x" + strings.Repeat("0", 60) + "5678"
			case strings.HasPrefix(callData, decryptCombineSelector):
				resp.Result = "0x" + strings.Repeat("0", 60) + "def0"
			case strings.HasPrefix(callData, revealShareSelector):
				resp.Result = "0x" + strings.Repeat("0", 60) + "2468"
			case strings.HasPrefix(callData, getNodeSelector):
				output, _ := registryABI.Methods["getNode"].Outputs.Pack(
					common.HexToAddress("0x4000000000000000000000000000000000000004"),
					big.NewInt(11),
					big.NewInt(12),
					uint8(1),
				)
				resp.Result = "0x" + hex.EncodeToString(output)
			case strings.HasPrefix(callData, getRoundSelector):
				// New lottery layout: 18 elements
				//  organizer, threshold, committeeSize, minValidContributions,
				//  lotteryAlphaBps, seedDelay, registrationDeadlineBlock,
				//  contributionDeadlineBlock, disclosureAllowed,
				//  status, nonce, seedBlock, seed, lotteryThreshold,
				//  claimedCount, contributionCount, partialDecryptionCount,
				//  revealedShareCount.
				type policyTuple struct {
					Threshold                 uint16 `json:"threshold"`
					CommitteeSize             uint16 `json:"committeeSize"`
					MinValidContributions     uint16 `json:"minValidContributions"`
					LotteryAlphaBps           uint16 `json:"lotteryAlphaBps"`
					SeedDelay                 uint16 `json:"seedDelay"`
					RegistrationDeadlineBlock uint64 `json:"registrationDeadlineBlock"`
					ContributionDeadlineBlock uint64 `json:"contributionDeadlineBlock"`
					FinalizeNotBeforeBlock    uint64 `json:"finalizeNotBeforeBlock"`
					DisclosureAllowed         bool   `json:"disclosureAllowed"`
				}
				type dpTuple struct {
					OwnerOnly          bool   `json:"ownerOnly"`
					MaxDecryptions     uint16 `json:"maxDecryptions"`
					NotBeforeBlock     uint64 `json:"notBeforeBlock"`
					NotBeforeTimestamp uint64 `json:"notBeforeTimestamp"`
					NotAfterBlock      uint64 `json:"notAfterBlock"`
					NotAfterTimestamp  uint64 `json:"notAfterTimestamp"`
				}
				output, _ := managerABI.Methods["getRound"].Outputs.Pack(
					common.HexToAddress("0x5000000000000000000000000000000000000005"),
					policyTuple{
						Threshold:                 2,
						CommitteeSize:             2,
						MinValidContributions:     2,
						LotteryAlphaBps:           20000,
						SeedDelay:                 4,
						RegistrationDeadlineBlock: 10,
						ContributionDeadlineBlock: 20,
						FinalizeNotBeforeBlock:    21,
						DisclosureAllowed:         false,
					},
					dpTuple{},                  // decryptionPolicy — all zero (no gate)
					uint8(3),                   // status
					uint64(1),                  // nonce
					uint64(8),                  // seedBlock
					common.HexToHash("0x5555"), // seed
					big.NewInt(1234567890),     // lotteryThreshold
					uint16(2),                  // claimedCount
					uint16(2),                  // contributionCount
					uint16(1),                  // partialDecryptionCount
					uint16(1),                  // revealedShareCount
					uint16(0),                  // ciphertextCount
				)
				resp.Result = "0x" + hex.EncodeToString(output)
			case strings.HasPrefix(callData, selectedParticipantsSelector):
				output, _ := managerABI.Methods["selectedParticipants"].Outputs.Pack([]common.Address{
					common.HexToAddress("0x5000000000000000000000000000000000000005"),
					common.HexToAddress("0x6000000000000000000000000000000000000006"),
				})
				resp.Result = "0x" + hex.EncodeToString(output)
			case strings.HasPrefix(callData, getCombinedDecryptionSelector):
				output, _ := managerABI.Methods["getCombinedDecryption"].Outputs.Pack(
					uint16(1),
					true,
					big.NewInt(39321), // 0x9999
				)
				resp.Result = "0x" + hex.EncodeToString(output)
			case strings.HasPrefix(callData, getRevealedShareSelector):
				output, _ := managerABI.Methods["getRevealedShare"].Outputs.Pack(
					common.HexToAddress("0x5000000000000000000000000000000000000005"),
					uint16(1),
					big.NewInt(7),
					common.HexToHash("0xaaaa"),
					true,
				)
				resp.Result = "0x" + hex.EncodeToString(output)
			default:
				resp.Error = &testRPCError{Code: -32601, Message: "unknown call data: " + callData}
			}
		default:
			resp.Error = &testRPCError{Code: -32601, Message: "method not found"}
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
}
