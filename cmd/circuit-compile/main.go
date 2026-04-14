package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	groth16_bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/backend/solidity"
	"github.com/consensys/gnark/constraint"
	flag "github.com/spf13/pflag"
	"github.com/vocdoni/davinci-dkg/circuits"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/decryptcombine"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	"github.com/vocdoni/davinci-dkg/circuits/partialdecrypt"
	"github.com/vocdoni/davinci-dkg/circuits/revealshare"
	"github.com/vocdoni/davinci-dkg/circuits/revealsubmit"
	"github.com/vocdoni/davinci-dkg/log"
	"github.com/vocdoni/davinci-dkg/prover"
)

type compiledArtifacts struct {
	CircuitHash      string `json:"circuit_hash"`
	ProvingKeyHash   string `json:"proving_key_hash"`
	VerifyingKeyHash string `json:"verifying_key_hash"`
	VerifierSolidity string `json:"verifier_solidity"`
}

type compileTarget struct {
	Name          string
	Artifacts     *circuits.CircuitArtifacts
	Compile       func() (constraint.ConstraintSystem, error)
	VerifierLabel string
}

func main() {
	var destination string
	var verifiersDir string
	var outputJSON string

	flag.StringVar(&destination, "destination", circuits.BaseDir, "destination folder for cached artifacts")
	flag.StringVar(&verifiersDir, "verifiers-dir", "solidity/src/verifiers", "destination folder for Solidity verifiers")
	flag.StringVar(&outputJSON, "output-json", "", "optional path to write the compile result JSON (bypasses stdout, which is polluted by gnark's own logger)")
	flag.Parse()

	log.Init("info", "stdout", nil)
	circuits.BaseDir = destination

	targets := []compileTarget{
		{
			Name:          "contribution",
			Artifacts:     contribution.Artifacts,
			Compile:       contribution.Compile,
			VerifierLabel: "contribution_vkey.sol",
		},
		{
			Name:          "finalize",
			Artifacts:     finalize.Artifacts,
			Compile:       finalize.Compile,
			VerifierLabel: "finalize_vkey.sol",
		},
		{
			Name:          "partialdecrypt",
			Artifacts:     partialdecrypt.Artifacts,
			Compile:       partialdecrypt.Compile,
			VerifierLabel: "partialdecrypt_vkey.sol",
		},
		{
			Name:          "decryptcombine",
			Artifacts:     decryptcombine.Artifacts,
			Compile:       decryptcombine.Compile,
			VerifierLabel: "decryptcombine_vkey.sol",
		},
		{
			Name:          "revealsubmit",
			Artifacts:     revealsubmit.Artifacts,
			Compile:       revealsubmit.Compile,
			VerifierLabel: "revealsubmit_vkey.sol",
		},
		{
			Name:          "revealshare",
			Artifacts:     revealshare.Artifacts,
			Compile:       revealshare.Compile,
			VerifierLabel: "revealshare_vkey.sol",
		},
	}

	results := map[string]compiledArtifacts{}
	for _, target := range targets {
		result, err := compileTargetArtifacts(target, destination, verifiersDir)
		if err != nil {
			log.Errorw(err, "compile target failed", "circuit", target.Name)
			os.Exit(1)
		}
		results[target.Name] = *result
	}

	encoded, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Errorw(err, "marshal artifact result")
		os.Exit(1)
	}
	// When --output-json is set write the result to that file so jq can
	// parse it cleanly. Otherwise fall back to stdout for interactive use
	// (the caller is responsible for stripping log lines).
	if outputJSON != "" {
		if err := os.WriteFile(outputJSON, encoded, 0o644); err != nil {
			log.Errorw(err, "write json", "path", outputJSON)
			os.Exit(1)
		}
	} else {
		fmt.Println(string(encoded))
	}
}

func compileTargetArtifacts(
	target compileTarget,
	destination string,
	verifiersDir string,
) (*compiledArtifacts, error) {
	log.Infow("compiling circuit definition", "circuit", target.Name)
	ccs, err := target.Compile()
	if err != nil {
		return nil, fmt.Errorf("compile %s: %w", target.Name, err)
	}

	runtime, err := target.Artifacts.Setup(ccs)
	if err != nil {
		return nil, fmt.Errorf("setup %s: %w", target.Name, err)
	}

	ccsArtifact, err := prover.WriteWriterArtifact(destination, ccs)
	if err != nil {
		return nil, fmt.Errorf("write %s ccs: %w", target.Name, err)
	}
	pkArtifact, err := prover.WriteWriterArtifact(destination, runtime.ProvingKey())
	if err != nil {
		return nil, fmt.Errorf("write %s proving key: %w", target.Name, err)
	}
	vkArtifact, err := prover.WriteWriterArtifact(destination, runtime.VerifyingKey())
	if err != nil {
		return nil, fmt.Errorf("write %s verifying key: %w", target.Name, err)
	}

	verifierPath, err := exportSolidityVerifier(
		filepath.Join(verifiersDir, target.VerifierLabel),
		pkArtifact.Hash,
		runtime.VerifyingKey(),
	)
	if err != nil {
		return nil, fmt.Errorf("export %s verifier: %w", target.Name, err)
	}

	return &compiledArtifacts{
		CircuitHash:      ccsArtifact.Hash,
		ProvingKeyHash:   pkArtifact.Hash,
		VerifyingKeyHash: vkArtifact.Hash,
		VerifierSolidity: verifierPath,
	}, nil
}

func exportSolidityVerifier(path string, provingKeyHash string, vk interface{}) (string, error) {
	solidityVK, ok := vk.(*groth16_bn254.VerifyingKey)
	if !ok {
		return "", fmt.Errorf("unexpected verifying key type %T", vk)
	}
	if err := solidityVK.Precompute(); err != nil {
		return "", fmt.Errorf("precompute verifying key: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return "", fmt.Errorf("create verifier dir: %w", err)
	}
	var buf bytes.Buffer
	if err := solidityVK.ExportSolidity(&buf, solidity.WithPragmaVersion("^0.8.28")); err != nil {
		return "", fmt.Errorf("export Solidity verifier: %w", err)
	}

	content := buf.String()
	marker := "contract Verifier"
	if idx := strings.Index(content, marker); idx >= 0 {
		content = content[:idx] +
			fmt.Sprintf("// provingKeyHash: %s\n", provingKeyHash) +
			content[idx:]
	}

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return "", fmt.Errorf("write verifier %s: %w", path, err)
	}
	return path, nil
}
