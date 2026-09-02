// sdk-test-fixture is a small helper binary used by the TypeScript SDK
// integration tests.  It connects to a running Anvil + deployer testnet,
// bootstraps the default Anvil node keys, and provides two actions:
//
//	--action=create  (default) creates a finalized single-participant DKG
//	                 epoch and writes JSON to stdout:
//	                   {"epochId":"0x…","collectivePublicKeyHash":"0x…","share":"<decimal>"}
//	                 The `share` is the polynomial share value held by
//	                 participant 1 (= the only contribution coefficient
//	                 used by the fixture), which the test passes back to
//	                 `--action=decrypt` so the helper can drive partial
//	                 decryption + combine over an SDK-submitted ciphertext.
//
//	--action=decrypt drives the threshold-decryption flow for a
//	                 ciphertext that the SDK already submitted on-chain:
//	                 builds the partial decryption proof, calls
//	                 submitPartialDecryption, releases the organizer share
//	                 and finally combineDecryption.
//	                 Required additional flags:
//	                   --epoch-id, --aid, --ciphertext-index, --share,
//	                   --org-secret
//	                 Outputs `{"ok":true}` on success.
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
	"github.com/vocdoni/davinci-dkg/log"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
)

type fixtureResult struct {
	EpochID                 string `json:"epochId"`
	CollectivePublicKeyHash string `json:"collectivePublicKeyHash"`
	Share                   string `json:"share"`
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

// fixtureShare is the polynomial share value held by participant 1 of the
// fixture epoch. CreateSDKTestFixture uses coefficients=[11] so f(1) = 11.
const fixtureShare int64 = 11

func main() {
	var rpcURL string
	var addressesFile string
	var action string
	var roundIDHex string
	var aidHex string
	var ciphertextIndex int
	var shareDec string
	var orgSecretHex string

	flag.StringVar(&rpcURL, "rpc-url", os.Getenv("DAVINCI_DKG_TEST_RPC_URL"),
		"RPC URL of the Anvil testnet")
	flag.StringVar(&addressesFile, "addresses-file", os.Getenv("DAVINCI_DKG_TEST_ADDRESSES"),
		"path to addresses.env file (as served by the deployer container)")
	flag.StringVar(&action, "action", "create",
		"action to perform: 'create' (default) or 'decrypt'")
	flag.StringVar(&roundIDHex, "epoch-id", "",
		"(decrypt) epoch id as a 0x-prefixed 12-byte hex string")
	flag.IntVar(&ciphertextIndex, "ciphertext-index", 0,
		"(decrypt) ciphertext index to combine (must be > 0)")
	flag.StringVar(&aidHex, "aid", "",
		"(decrypt) application id as a 0x-prefixed 32-byte hex string")
	flag.StringVar(&shareDec, "share", "",
		"(decrypt) participant 1's polynomial share value, decimal")
	flag.StringVar(&orgSecretHex, "org-secret", "",
		"(decrypt) organizer secret the application was registered with, hex scalar")
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	services, cleanup, err := helpers.NewTestServicesFromExternal(ctx, rpcURL, addressesContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: setup services: %v\n", err)
		os.Exit(1)
	}
	defer cleanup()

	switch action {
	case "create":
		result, err := helpers.CreateSDKTestFixture(ctx, services)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: create fixture: %v\n", err)
			os.Exit(1)
		}
		out := fixtureResult{
			EpochID:                 fmt.Sprintf("0x%x", result.EpochID),
			CollectivePublicKeyHash: fmt.Sprintf("0x%x", result.CollectivePublicKeyHash),
			Share:                   big.NewInt(fixtureShare).String(),
		}
		encoded, err := json.Marshal(out)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: marshal result: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(encoded))

	case "decrypt":
		if roundIDHex == "" {
			fmt.Fprintln(os.Stderr, "error: --epoch-id is required for decrypt")
			os.Exit(1)
		}
		if ciphertextIndex <= 0 || ciphertextIndex > 0xffff {
			fmt.Fprintln(os.Stderr, "error: --ciphertext-index must be in (0, 65535]")
			os.Exit(1)
		}
		if shareDec == "" {
			fmt.Fprintln(os.Stderr, "error: --share is required for decrypt")
			os.Exit(1)
		}
		share, ok := new(big.Int).SetString(shareDec, 10)
		if !ok {
			fmt.Fprintf(os.Stderr, "error: --share %q is not a valid decimal\n", shareDec)
			os.Exit(1)
		}
		epochID := mustEpochID(roundIDHex)
		aid := mustAid(aidHex)
		skOrg := mustOrganizerSecret(orgSecretHex)

		if err := helpers.CombineSingleParticipantDecryption(
			ctx, services, epochID, aid, uint16(ciphertextIndex), share, skOrg,
		); err != nil {
			fmt.Fprintf(os.Stderr, "error: combine decryption: %v\n", err)
			os.Exit(1)
		}
		encoded, err := json.Marshal(decryptResult{OK: true})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: marshal result: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(encoded))

	case "prepare-combine":
		if roundIDHex == "" {
			fmt.Fprintln(os.Stderr, "error: --epoch-id is required for prepare-combine")
			os.Exit(1)
		}
		if ciphertextIndex <= 0 || ciphertextIndex > 0xffff {
			fmt.Fprintln(os.Stderr, "error: --ciphertext-index must be in (0, 65535]")
			os.Exit(1)
		}
		if shareDec == "" {
			fmt.Fprintln(os.Stderr, "error: --share is required for prepare-combine")
			os.Exit(1)
		}
		share, ok := new(big.Int).SetString(shareDec, 10)
		if !ok {
			fmt.Fprintf(os.Stderr, "error: --share %q is not a valid decimal\n", shareDec)
			os.Exit(1)
		}
		epochID := mustEpochID(roundIDHex)
		aid := mustAid(aidHex)
		skOrg := mustOrganizerSecret(orgSecretHex)

		payload, err := helpers.PrepareSingleParticipantCombinePayload(
			ctx, services, epochID, aid, uint16(ciphertextIndex), share, skOrg,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: prepare combine: %v\n", err)
			os.Exit(1)
		}
		out := prepareCombineResult{
			CombineHash: "0x" + hex.EncodeToString(payload.CombineHash[:]),
			Plaintext:   payload.Plaintext.String(),
			Transcript:  "0x" + hex.EncodeToString(payload.Transcript),
			Proof:       "0x" + hex.EncodeToString(payload.Proof),
			Input:       "0x" + hex.EncodeToString(payload.Input),
		}
		encoded, err := json.Marshal(out)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: marshal result: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(encoded))

	default:
		fmt.Fprintf(os.Stderr, "error: unknown --action %q (must be 'create' or 'decrypt')\n", action)
		os.Exit(1)
	}
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

// mustOrganizerSecret parses the organizer scalar the application was
// registered with, or exits. Without it the ciphertext cannot be decrypted:
// the committee alone only recovers sk_ep·C1.
func mustOrganizerSecret(value string) *big.Int {
	sk, ok := new(big.Int).SetString(strings.TrimPrefix(value, "0x"), 16)
	if !ok || sk.Sign() <= 0 {
		fmt.Fprintf(os.Stderr, "error: --org-secret must be a positive hex scalar, got %q\n", value)
		os.Exit(1)
	}
	return sk
}
