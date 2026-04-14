// sdk-test-fixture is a small helper binary used by the TypeScript SDK
// integration tests.  It connects to a running Anvil + deployer testnet,
// bootstraps the default Anvil node keys, creates a finalized
// single-participant DKG round, and writes a JSON result to stdout:
//
//	{"roundId":"0x…","collectivePublicKeyHash":"0x…"}
//
// The TypeScript tests use this output to verify monitoring and read
// operations against a fully-finalized round state without needing to
// generate Groth16 proofs in TypeScript.
//
// Flags (all can be provided as CLI flags or their env-var equivalents):
//
//	--rpc-url        RPC_URL env var or DAVINCI_DKG_TEST_RPC_URL
//	--addresses-file path to addresses.env served by the deployer container
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/vocdoni/davinci-dkg/log"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
)

type fixtureResult struct {
	RoundID                 string `json:"roundId"`
	CollectivePublicKeyHash string `json:"collectivePublicKeyHash"`
}

func main() {
	var rpcURL string
	var addressesFile string

	flag.StringVar(&rpcURL, "rpc-url", os.Getenv("DAVINCI_DKG_TEST_RPC_URL"),
		"RPC URL of the Anvil testnet")
	flag.StringVar(&addressesFile, "addresses-file", os.Getenv("DAVINCI_DKG_TEST_ADDRESSES"),
		"path to addresses.env file (as served by the deployer container)")
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

	result, err := helpers.CreateSDKTestFixture(ctx, services)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: create fixture: %v\n", err)
		os.Exit(1)
	}

	out := fixtureResult{
		RoundID:                 fmt.Sprintf("0x%x", result.RoundID),
		CollectivePublicKeyHash: fmt.Sprintf("0x%x", result.CollectivePublicKeyHash),
	}
	encoded, err := json.Marshal(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: marshal result: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(encoded))
}
