package main

import (
	"bytes"
	"context"
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
	s3Config := NewDefaultS3Config()

	flag.StringVar(&destination, "destination", circuits.BaseDir, "destination folder for cached artifacts")
	flag.StringVar(&verifiersDir, "verifiers-dir", "solidity/src/verifiers", "destination folder for Solidity verifiers")
	flag.StringVar(&outputJSON, "output-json", "", "optional path to write the compile result JSON (bypasses stdout, which is polluted by gnark's own logger)")
	flag.BoolVar(&s3Config.Enabled, "s3.enabled", false, "upload compiled artifacts to DigitalOcean Spaces / S3")
	flag.StringVar(&s3Config.AccessKey, "s3.access-key", "", "S3 / DO Spaces access key")
	flag.StringVar(&s3Config.SecretKey, "s3.secret-key", "", "S3 / DO Spaces secret key")
	flag.StringVar(&s3Config.Space, "s3.space", "circuits", "DO Spaces bucket name")
	flag.StringVar(&s3Config.Bucket, "s3.bucket", "dev", "folder within the bucket (release channel)")
	flag.Parse()

	log.Init("info", "stdout", nil)
	circuits.BaseDir = destination

	if s3Config.Enabled {
		if err := TestS3Connection(context.Background(), s3Config); err != nil {
			log.Errorw(err, "S3 connection test failed")
			os.Exit(1)
		}
	}

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

	if s3Config.Enabled {
		if err := uploadArtifacts(context.Background(), s3Config, destination, results); err != nil {
			log.Errorw(err, "upload artifacts failed")
			os.Exit(1)
		}
	}
}

// uploadArtifacts uploads all compiled circuit artifacts to S3/DO Spaces.
// Each file is stored under its content-hash with the appropriate extension
// (.ccs / .pk / .vk) matching the URL scheme in config/circuit_artifacts.go.
func uploadArtifacts(ctx context.Context, s3Config *S3Config, destination string, results map[string]compiledArtifacts) error {
	uploader, err := NewS3Uploader(s3Config)
	if err != nil {
		return fmt.Errorf("create uploader: %w", err)
	}

	type upload struct {
		localPath  string
		remoteName string
	}
	var uploads []upload
	for name, a := range results {
		_ = name
		uploads = append(uploads,
			upload{filepath.Join(destination, a.CircuitHash), a.CircuitHash + ".ccs"},
			upload{filepath.Join(destination, a.ProvingKeyHash), a.ProvingKeyHash + ".pk"},
			upload{filepath.Join(destination, a.VerifyingKeyHash), a.VerifyingKeyHash + ".vk"},
		)
	}

	var uploadedKeys []string
	for _, u := range uploads {
		key, err := uploader.UploadFileAs(ctx, u.localPath, u.remoteName)
		if err != nil {
			return fmt.Errorf("upload %s: %w", u.remoteName, err)
		}
		uploadedKeys = append(uploadedKeys, key)
	}

	if err := uploader.SetPublicACL(ctx, uploadedKeys); err != nil {
		return fmt.Errorf("set public ACL: %w", err)
	}

	log.Infow("all artifacts uploaded and made public", "count", len(uploadedKeys),
		"space", s3Config.Space, "bucket", s3Config.Bucket)
	return nil
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
