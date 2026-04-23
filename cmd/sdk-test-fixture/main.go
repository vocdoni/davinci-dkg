// sdk-test-fixture is a small helper binary used by the TypeScript SDK
// integration tests.  It connects to a running Anvil + deployer testnet,
// bootstraps the default Anvil node keys, and provides two actions:
//
//	--action=create  (default) creates a finalized single-participant DKG
//	                 round and writes JSON to stdout:
//	                   {"roundId":"0x…","collectivePublicKeyHash":"0x…","share":"<decimal>"}
//	                 The `share` is the polynomial share value held by
//	                 participant 1 (= the only contribution coefficient
//	                 used by the fixture), which the test passes back to
//	                 `--action=decrypt` so the helper can drive partial
//	                 decryption + combine over an SDK-submitted ciphertext.
//
//	--action=decrypt drives the threshold-decryption flow for a
//	                 ciphertext that the SDK already submitted on-chain:
//	                 builds the partial decryption proof, calls
//	                 submitPartialDecryption, then combineDecryption.
//	                 Required additional flags:
//	                   --round-id, --ciphertext-index, --share
//	                 Outputs `{"ok":true}` on success.
//
// The TypeScript tests use these together to verify the full round-trip
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
	RoundID                 string `json:"roundId"`
	CollectivePublicKeyHash string `json:"collectivePublicKeyHash"`
	Share                   string `json:"share"`
}

type decryptResult struct {
	OK bool `json:"ok"`
}

// fixtureShare is the polynomial share value held by participant 1 of the
// fixture round. CreateSDKTestFixture uses coefficients=[11] so f(1) = 11.
const fixtureShare int64 = 11

func main() {
	var rpcURL string
	var addressesFile string
	var action string
	var roundIDHex string
	var ciphertextIndex int
	var shareDec string

	flag.StringVar(&rpcURL, "rpc-url", os.Getenv("DAVINCI_DKG_TEST_RPC_URL"),
		"RPC URL of the Anvil testnet")
	flag.StringVar(&addressesFile, "addresses-file", os.Getenv("DAVINCI_DKG_TEST_ADDRESSES"),
		"path to addresses.env file (as served by the deployer container)")
	flag.StringVar(&action, "action", "create",
		"action to perform: 'create' (default) or 'decrypt'")
	flag.StringVar(&roundIDHex, "round-id", "",
		"(decrypt) round id as a 0x-prefixed 12-byte hex string")
	flag.IntVar(&ciphertextIndex, "ciphertext-index", 0,
		"(decrypt) ciphertext index to combine (must be > 0)")
	flag.StringVar(&shareDec, "share", "",
		"(decrypt) participant 1's polynomial share value, decimal")
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
			RoundID:                 fmt.Sprintf("0x%x", result.RoundID),
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
			fmt.Fprintln(os.Stderr, "error: --round-id is required for decrypt")
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
		raw, err := hex.DecodeString(strings.TrimPrefix(roundIDHex, "0x"))
		if err != nil || len(raw) != 12 {
			fmt.Fprintf(os.Stderr, "error: --round-id must be 0x-prefixed 12-byte hex, got %q\n", roundIDHex)
			os.Exit(1)
		}
		var roundID [12]byte
		copy(roundID[:], raw)

		if err := helpers.CombineSingleParticipantDecryption(ctx, services, roundID, uint16(ciphertextIndex), share); err != nil {
			fmt.Fprintf(os.Stderr, "error: combine decryption: %v\n", err)
			os.Exit(1)
		}
		encoded, err := json.Marshal(decryptResult{OK: true})
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
