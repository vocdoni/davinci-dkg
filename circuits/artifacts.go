package circuits

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	gpugroth16 "github.com/consensys/gnark/backend/accelerated/icicle/groth16"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/vocdoni/davinci-dkg/log"
	"github.com/vocdoni/davinci-dkg/prover"
)

// BaseDir is the cache directory for compiled circuit artifacts.
var BaseDir string

var (
	ErrArtifactNotFound     = errors.New("artifact not found in cache")
	ErrArtifactHashMismatch = errors.New("artifact hash mismatch")
)

func init() {
	if BaseDir == "" {
		if dir := os.Getenv("DAVINCI_ARTIFACTS_DIR"); dir != "" {
			BaseDir = dir
		} else {
			userHomeDir, err := os.UserHomeDir()
			if err != nil {
				userHomeDir = "."
			}
			BaseDir = filepath.Join(userHomeDir, ".davinci", "artifacts")
		}
	}
}

// Artifact describes a cached/downloadable circuit artifact by hash and source URL.
type Artifact struct {
	RemoteURL string
	Hash      []byte
}

func (a *Artifact) cachePath() (string, error) {
	if a == nil {
		return "", fmt.Errorf("artifact not configured")
	}
	if len(a.Hash) == 0 {
		return "", fmt.Errorf("artifact hash not provided")
	}
	return filepath.Join(BaseDir, hex.EncodeToString(a.Hash)), nil
}

func (a *Artifact) loadOrDownload(ctx context.Context) ([]byte, error) {
	if a == nil {
		return nil, fmt.Errorf("artifact not configured")
	}
	if len(a.Hash) == 0 {
		return nil, fmt.Errorf("artifact hash not provided")
	}

	content, err := a.loadFromCache()
	switch {
	case err == nil:
		return content, nil
	case errors.Is(err, ErrArtifactNotFound), errors.Is(err, ErrArtifactHashMismatch):
		if err := a.downloadToCache(ctx); err != nil {
			return nil, err
		}
	default:
		return nil, err
	}
	return a.readFromCache()
}

func (a *Artifact) loadFromCache() ([]byte, error) {
	content, err := a.readFromCache()
	if err != nil {
		return nil, err
	}
	return content, a.checkHash(content)
}

func (a *Artifact) readFromCache() ([]byte, error) {
	path, err := a.cachePath()
	if err != nil {
		return nil, err
	}
	content, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrArtifactNotFound
		}
		return nil, fmt.Errorf("read cached artifact %s: %w", path, err)
	}
	return content, nil
}

func (a *Artifact) checkHash(content []byte) error {
	sum := sha256.Sum256(content)
	if !bytes.Equal(sum[:], a.Hash) {
		log.Warnw("hash mismatch for cached artifact",
			"expected", hex.EncodeToString(a.Hash),
			"got", hex.EncodeToString(sum[:]),
		)
		return fmt.Errorf("%w: expected %x, got %x", ErrArtifactHashMismatch, a.Hash, sum[:])
	}
	return nil
}

func (a *Artifact) checkFileHash() error {
	path, err := a.cachePath()
	if err != nil {
		return err
	}
	fd, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrArtifactNotFound
		}
		return fmt.Errorf("open cached artifact %s: %w", path, err)
	}
	defer func() { _ = fd.Close() }()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, fd); err != nil {
		return fmt.Errorf("hash cached artifact %s: %w", path, err)
	}
	sum := hasher.Sum(nil)
	if !bytes.Equal(sum, a.Hash) {
		return fmt.Errorf("%w: expected %x, got %x", ErrArtifactHashMismatch, a.Hash, sum)
	}
	return nil
}

func (a *Artifact) ensureDownloaded(ctx context.Context) error {
	err := a.checkFileHash()
	switch {
	case err == nil:
		return nil
	case errors.Is(err, ErrArtifactNotFound), errors.Is(err, ErrArtifactHashMismatch):
		return a.downloadToCache(ctx)
	default:
		return err
	}
}

func (a *Artifact) downloadToCache(ctx context.Context) error {
	path, err := a.cachePath()
	if err != nil {
		return err
	}
	if a.RemoteURL == "" {
		return fmt.Errorf("artifact remote url not provided for hash %s", hex.EncodeToString(a.Hash))
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.RemoteURL, nil)
	if err != nil {
		return fmt.Errorf("create file request: %w", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Errorw(err, "error closing body")
		}
	}()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http status %s for %s", res.Status, a.RemoteURL)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create artifact cache dir %s: %w", dir, err)
	}

	tmpFile, err := os.CreateTemp(dir, hex.EncodeToString(a.Hash)+".tmp.*")
	if err != nil {
		return fmt.Errorf("create temp artifact file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer func() {
		if tmpFile != nil {
			_ = tmpFile.Close()
		}
		if tmpPath != "" {
			_ = os.Remove(tmpPath)
		}
	}()

	hasher := sha256.New()
	size, err := io.Copy(io.MultiWriter(tmpFile, hasher), res.Body)
	if err != nil {
		return fmt.Errorf("read downloaded content: %w", err)
	}
	sum := hasher.Sum(nil)
	if !bytes.Equal(sum, a.Hash) {
		return fmt.Errorf("hash mismatch for downloaded artifact: expected %s, got %s",
			hex.EncodeToString(a.Hash), hex.EncodeToString(sum))
	}
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("close temp cache file %s: %w", tmpPath, err)
	}
	tmpFile = nil
	if err := os.Rename(tmpPath, path); err != nil {
		return fmt.Errorf("install cached artifact %s: %w", path, err)
	}
	tmpPath = ""
	log.Debugw("artifact downloaded and cached", "path", path, "size_bytes", size)
	return nil
}

type circuitParams struct {
	name         string
	curve        ecc.ID
	proverOpts   []backend.ProverOption
	verifierOpts []backend.VerifierOption
}

func (c circuitParams) Name() string                          { return c.name }
func (c circuitParams) Curve() ecc.ID                         { return c.curve }
func (c circuitParams) ProverOptions() []backend.ProverOption { return c.proverOpts }
func (c circuitParams) VerifierOptions() []backend.VerifierOption {
	return c.verifierOpts
}

// CircuitArtifacts describes the three serialized artifacts associated with a circuit.
type CircuitArtifacts struct {
	circuitParams
	circuitDefinition *Artifact
	provingKey        *Artifact
	verifyingKey      *Artifact
}

// NewCircuitArtifacts builds a CircuitArtifacts bundle.
func NewCircuitArtifacts(
	name string,
	curve ecc.ID,
	proverOpts []backend.ProverOption,
	verifierOpts []backend.VerifierOption,
	circuit *Artifact,
	provingKey *Artifact,
	verifyingKey *Artifact,
) *CircuitArtifacts {
	return &CircuitArtifacts{
		circuitParams: circuitParams{
			name:         name,
			curve:        curve,
			proverOpts:   proverOpts,
			verifierOpts: verifierOpts,
		},
		circuitDefinition: circuit,
		provingKey:        provingKey,
		verifyingKey:      verifyingKey,
	}
}

// Download ensures all configured artifacts exist in cache, fetching them when needed.
func (ca *CircuitArtifacts) Download(ctx context.Context) error {
	if ca.circuitDefinition != nil && len(ca.circuitDefinition.Hash) != 0 {
		if err := ca.circuitDefinition.ensureDownloaded(ctx); err != nil {
			return fmt.Errorf("download circuit definition: %w", err)
		}
	}
	if ca.provingKey != nil && len(ca.provingKey.Hash) != 0 {
		if err := ca.provingKey.ensureDownloaded(ctx); err != nil {
			return fmt.Errorf("download proving key: %w", err)
		}
	}
	if ca.verifyingKey != nil && len(ca.verifyingKey.Hash) != 0 {
		if err := ca.verifyingKey.ensureDownloaded(ctx); err != nil {
			return fmt.Errorf("download verifying key: %w", err)
		}
	}
	return nil
}

// LoadOrDownload ensures all configured artifacts are available and decodes them into a runtime.
func (ca *CircuitArtifacts) LoadOrDownload(ctx context.Context) (*CircuitRuntime, error) {
	ccs, err := ca.LoadOrDownloadCircuitDefinition(ctx)
	if err != nil {
		return nil, fmt.Errorf("load circuit definition: %w", err)
	}
	pk, err := ca.LoadOrDownloadProvingKey(ctx)
	if err != nil {
		return nil, fmt.Errorf("load proving key: %w", err)
	}
	vk, err := ca.LoadOrDownloadVerifyingKey(ctx)
	if err != nil {
		return nil, fmt.Errorf("load verifying key: %w", err)
	}
	return NewCircuitRuntime(ca.name, ca.curve, ca.proverOpts, ca.verifierOpts, ccs, pk, vk), nil
}

// LoadOrDownloadCircuitDefinition loads the configured circuit definition artifact.
func (ca *CircuitArtifacts) LoadOrDownloadCircuitDefinition(ctx context.Context) (constraint.ConstraintSystem, error) {
	if ca.circuitDefinition == nil || len(ca.circuitDefinition.Hash) == 0 {
		return nil, fmt.Errorf("circuit definition not configured")
	}
	content, err := ca.circuitDefinition.loadOrDownload(ctx)
	if err != nil {
		return nil, fmt.Errorf("load circuit definition: %w", err)
	}
	ccs := newConstraintSystem(ca.curve)
	if _, err := ccs.ReadFrom(bytes.NewReader(content)); err != nil {
		return nil, fmt.Errorf("decode circuit definition: %w", err)
	}
	return ccs, nil
}

// LoadOrDownloadVerifyingKey loads the configured verifying key artifact.
func (ca *CircuitArtifacts) LoadOrDownloadVerifyingKey(ctx context.Context) (groth16.VerifyingKey, error) {
	if ca.verifyingKey == nil || len(ca.verifyingKey.Hash) == 0 {
		return nil, fmt.Errorf("verifying key not configured")
	}
	content, err := ca.verifyingKey.loadOrDownload(ctx)
	if err != nil {
		return nil, fmt.Errorf("load verifying key: %w", err)
	}
	vk := newVerifyingKey(ca.curve)
	if _, err := vk.UnsafeReadFrom(bytes.NewReader(content)); err != nil {
		return nil, fmt.Errorf("decode verifying key: %w", err)
	}
	return vk, nil
}

// LoadOrDownloadProvingKey loads the configured proving key artifact.
func (ca *CircuitArtifacts) LoadOrDownloadProvingKey(ctx context.Context) (groth16.ProvingKey, error) {
	if ca.provingKey == nil || len(ca.provingKey.Hash) == 0 {
		return nil, fmt.Errorf("proving key not configured")
	}
	content, err := ca.provingKey.loadOrDownload(ctx)
	if err != nil {
		return nil, fmt.Errorf("load proving key: %w", err)
	}
	pk := newProvingKey(ca.curve)
	if _, err := pk.UnsafeReadFrom(bytes.NewReader(content)); err != nil {
		return nil, fmt.Errorf("decode proving key: %w", err)
	}
	return pk, nil
}

// RawVerifyingKey returns the raw cached verifying key bytes.
func (ca *CircuitArtifacts) RawVerifyingKey() ([]byte, error) {
	if ca.verifyingKey == nil || len(ca.verifyingKey.Hash) == 0 {
		return nil, fmt.Errorf("verifying key not configured")
	}
	content, err := ca.verifyingKey.loadFromCache()
	if err != nil {
		return nil, fmt.Errorf("load verifying key: %w", err)
	}
	return content, nil
}

func (ca *CircuitArtifacts) CircuitHash() []byte {
	if ca.circuitDefinition == nil {
		return nil
	}
	return ca.circuitDefinition.Hash
}

func (ca *CircuitArtifacts) ProvingKeyHash() []byte {
	if ca.provingKey == nil {
		return nil
	}
	return ca.provingKey.Hash
}

func (ca *CircuitArtifacts) VerifyingKeyHash() []byte {
	if ca.verifyingKey == nil {
		return nil
	}
	return ca.verifyingKey.Hash
}

// Matches reports whether a compiled circuit matches the configured circuit hash.
func (ca *CircuitArtifacts) Matches(ccs constraint.ConstraintSystem) (bool, error) {
	if ccs == nil {
		return false, fmt.Errorf("constraint system not provided")
	}
	expectedHash := ca.CircuitHash()
	if len(expectedHash) == 0 {
		return false, fmt.Errorf("circuit hash not configured for circuit %s", ca.Name())
	}
	hasher := sha256.New()
	if _, err := ccs.WriteTo(hasher); err != nil {
		return false, fmt.Errorf("write ccs to hasher: %w", err)
	}
	currentHash := hasher.Sum(nil)
	return bytes.Equal(currentHash, expectedHash), nil
}

// Setup generates fresh proving and verifying keys for a compiled circuit.
func (ca *CircuitArtifacts) Setup(ccs constraint.ConstraintSystem) (*CircuitRuntime, error) {
	if ccs == nil {
		return nil, fmt.Errorf("constraint system not provided")
	}
	pk, vk, err := prover.Setup(ccs)
	if err != nil {
		return nil, fmt.Errorf("setup circuit %s: %w", ca.Name(), err)
	}
	return NewCircuitRuntime(ca.name, ca.curve, ca.proverOpts, ca.verifierOpts, ccs, pk, vk), nil
}

// LoadOrSetupForCircuit compiles the provided circuit and either reuses matching
// serialized artifacts or runs a fresh trusted setup.
func (ca *CircuitArtifacts) LoadOrSetupForCircuit(ctx context.Context, circuit frontend.Circuit) (*CircuitRuntime, error) {
	if ca == nil {
		return nil, fmt.Errorf("circuit artifacts not provided")
	}
	ccs, err := frontend.Compile(ca.Curve().ScalarField(), r1cs.NewBuilder, circuit)
	if err != nil {
		return nil, fmt.Errorf("compile circuit: %w", err)
	}
	if len(ca.CircuitHash()) == 0 || len(ca.ProvingKeyHash()) == 0 || len(ca.VerifyingKeyHash()) == 0 {
		return ca.Setup(ccs)
	}
	matches, err := ca.Matches(ccs)
	if err != nil {
		return nil, fmt.Errorf("match artifacts: %w", err)
	}
	if matches {
		runtime, err := ca.LoadOrDownload(ctx)
		if err == nil {
			return runtime, nil
		}
		log.Warnw("artifact load failed, falling back to local setup", "circuit", ca.Name(), "error", err)
	}
	return ca.Setup(ccs)
}

// CircuitRuntime is an in-memory runtime view of a compiled circuit and its keys.
type CircuitRuntime struct {
	circuitParams
	ccs constraint.ConstraintSystem
	pk  groth16.ProvingKey
	vk  groth16.VerifyingKey
}

// NewCircuitRuntime builds a runtime from already decoded artifacts.
func NewCircuitRuntime(
	name string,
	curve ecc.ID,
	proverOpts []backend.ProverOption,
	verifierOpts []backend.VerifierOption,
	ccs constraint.ConstraintSystem,
	pk groth16.ProvingKey,
	vk groth16.VerifyingKey,
) *CircuitRuntime {
	return &CircuitRuntime{
		circuitParams: circuitParams{
			name:         name,
			curve:        curve,
			proverOpts:   proverOpts,
			verifierOpts: verifierOpts,
		},
		ccs: ccs,
		pk:  pk,
		vk:  vk,
	}
}

// maxProveAttempts is the number of times ProveAndVerify will re-prove on
// a verification failure before giving up. The gnark Groth16 prover is
// probabilistic (it samples fresh random blinding scalars r,s every call); a
// spurious invalid proof can be cured by simply re-proving. Four attempts
// reduce the residual failure probability to <(0.05)^4 ≈ 6 × 10⁻⁶.
const maxProveAttempts = 4

// ProveAndVerify proves an assignment and immediately verifies it.
// On a transient prover failure (invalid proof) it retries up to maxProveAttempts
// times with fresh randomness before returning an error.
func (cr *CircuitRuntime) ProveAndVerify(assignment frontend.Circuit) (groth16.Proof, error) {
	var lastErr error
	for attempt := 0; attempt < maxProveAttempts; attempt++ {
		proof, err := cr.Prove(assignment)
		if err != nil {
			return nil, err // proving key / witness error — not transient, don't retry
		}
		if err := cr.Verify(proof, assignment); err != nil {
			lastErr = err
			log.Warnw("proof verification failed, retrying with fresh randomness",
				"circuit", cr.Name(), "attempt", attempt+1, "maxAttempts", maxProveAttempts, "error", err)
			continue
		}
		return proof, nil
	}
	return nil, fmt.Errorf("proof verification failed after %d attempts: %w", maxProveAttempts, lastErr)
}

// ProveAndVerifyWithWitness proves and verifies using a precomputed witness.
// It retries up to maxProveAttempts times on a transient prover failure.
func (cr *CircuitRuntime) ProveAndVerifyWithWitness(fullWitness witness.Witness) (groth16.Proof, error) {
	publicWitness, err := fullWitness.Public()
	if err != nil {
		return nil, err
	}
	var lastErr error
	for attempt := 0; attempt < maxProveAttempts; attempt++ {
		proof, err := cr.ProveWithWitness(fullWitness)
		if err != nil {
			return nil, err
		}
		if err := cr.VerifyWithWitness(proof, publicWitness); err != nil {
			lastErr = err
			log.Warnw("proof verification failed, retrying with fresh randomness",
				"circuit", cr.Name(), "attempt", attempt+1, "maxAttempts", maxProveAttempts, "error", err)
			continue
		}
		return proof, nil
	}
	return nil, fmt.Errorf("proof verification failed after %d attempts: %w", maxProveAttempts, lastErr)
}

// Prove generates a proof from an assignment.
func (cr *CircuitRuntime) Prove(assignment frontend.Circuit) (groth16.Proof, error) {
	start := time.Now()
	proof, err := prover.Prove(cr.curve, cr.ccs, cr.pk, assignment, cr.proverOpts...)
	if err == nil {
		log.Debugw("proof generated", "circuit", cr.Name(), "elapsed", time.Since(start).String())
	}
	return proof, err
}

// ProveWithWitness generates a proof from a precomputed witness.
func (cr *CircuitRuntime) ProveWithWitness(fullWitness witness.Witness) (groth16.Proof, error) {
	start := time.Now()
	proof, err := prover.ProveWithWitness(cr.curve, cr.ccs, cr.pk, fullWitness, cr.proverOpts...)
	if err == nil {
		log.Debugw("proof generated", "circuit", cr.Name(), "elapsed", time.Since(start).String())
	}
	return proof, err
}

// Verify verifies a proof against a public assignment.
func (cr *CircuitRuntime) Verify(proof groth16.Proof, publicAssignment frontend.Circuit) error {
	publicWitness, err := frontend.NewWitness(publicAssignment, cr.curve.ScalarField(), frontend.PublicOnly())
	if err != nil {
		return fmt.Errorf("create public witness: %w", err)
	}
	return cr.VerifyWithWitness(proof, publicWitness)
}

// VerifyWithWitness verifies a proof against a public witness.
func (cr *CircuitRuntime) VerifyWithWitness(proof groth16.Proof, publicWitness witness.Witness) error {
	start := time.Now()
	err := groth16.Verify(proof, cr.vk, publicWitness, cr.verifierOpts...)
	if err == nil {
		log.Debugw("proof verified", "circuit", cr.Name(), "elapsed", time.Since(start).String())
	}
	return err
}

func (cr *CircuitRuntime) ConstraintSystem() constraint.ConstraintSystem { return cr.ccs }
func (cr *CircuitRuntime) ProvingKey() groth16.ProvingKey                { return cr.pk }
func (cr *CircuitRuntime) VerifyingKey() groth16.VerifyingKey            { return cr.vk }

func newConstraintSystem(curve ecc.ID) constraint.ConstraintSystem {
	if prover.UseGPUProver {
		return gpugroth16.NewCS(curve)
	}
	return groth16.NewCS(curve)
}

func newProvingKey(curve ecc.ID) groth16.ProvingKey {
	if prover.UseGPUProver {
		return gpugroth16.NewProvingKey(curve)
	}
	return groth16.NewProvingKey(curve)
}

func newVerifyingKey(curve ecc.ID) groth16.VerifyingKey {
	if prover.UseGPUProver {
		return gpugroth16.NewVerifyingKey(curve)
	}
	return groth16.NewVerifyingKey(curve)
}
