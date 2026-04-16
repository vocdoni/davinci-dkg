// dkg-runner is the testnet scenario orchestrator.
//
// It creates a DKG round, waits for N DKG nodes to signal readiness and submit
// contributions, finalises the round, encrypts a test message under the
// collective public key, waits for all nodes to submit partial decryptions,
// combines them, and asserts the recovered plaintext equals the original.
//
// All parameters are read from environment variables (DKG_RUNNER_ prefix) or
// the flags below.  The runner is the ONLY actor that submits the finalize and
// combine transactions; everything else is done by the participating nodes.
//
// Usage:
//
//	dkg-runner --rpc http://anvil:8545 --manager 0x... \
//	           --privkey 0x... --nodes 3 --threshold 2
package main

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/vocdoni/davinci-dkg/circuits"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/decryptcombine"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	"github.com/vocdoni/davinci-dkg/config"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/log"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	nodetypes "github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
)

// CiphertextFile is read by node daemons from sharedDir/<roundHex>.json.
type CiphertextFile struct {
	CiphertextIndex uint16 `json:"ciphertext_index"`
	C1X             string `json:"c1x"`
	C1Y             string `json:"c1y"`
}

type cfg struct {
	RPC                        string
	PrivKey                    string
	Network                    string
	Manager                    string
	Nodes                      int
	Threshold                  int
	SharedDir                  string
	ArtifactsDir               string
	LogLevel                   string
	WaitReadiness              time.Duration
	WaitContrib                time.Duration
	WaitDecrypt                time.Duration
	ReadinessDeadlineBlocks    int
	ContributionDeadlineBlocks int
	DisclosureAllowed          bool
}

func loadCfg() (*cfg, error) {
	fs := flag.NewFlagSet("dkg-runner", flag.ContinueOnError)
	fs.String("rpc", "http://127.0.0.1:8545", "Ethereum RPC endpoint")
	fs.String("privkey", "", "organiser hex private key")
	fs.String("network", "", "well-known network preset (e.g. sepolia, sep); sets manager address automatically")
	fs.String("manager", "", "DKGManager contract address (optional when --network is set)")
	fs.Int("nodes", 3, "number of DKG nodes to wait for")
	fs.Int("threshold", 2, "decryption threshold")
	fs.String("shared-dir", "/shared", "directory written with ciphertext files")
	fs.String("artifacts-dir", "", "circuit artifact cache directory")
	fs.String("log-level", "info", "log level")
	fs.Duration("wait-readiness", 4*time.Minute, "timeout for readiness phase")
	fs.Duration("wait-contrib", 20*time.Minute, "timeout for contribution phase (large circuits need time to cold-load PK)")
	fs.Duration("wait-decrypt", 20*time.Minute, "timeout for partial-decryption phase")
	fs.Int("readiness-deadline-blocks", 30, "block offset from current head for readiness deadline")
	fs.Int("contribution-deadline-blocks", 600, "block offset from current head for contribution deadline (~20 min at 2s blocks)")
	fs.Bool("disclosure-allowed", false, "enable the reveal-share disclosure/reconstruction phase on the round")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return nil, err
	}
	v := viper.New()
	v.SetEnvPrefix("DKG_RUNNER")
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()
	if err := v.BindPFlags(fs); err != nil {
		return nil, err
	}
	c := &cfg{
		RPC:                        v.GetString("rpc"),
		PrivKey:                    v.GetString("privkey"),
		Network:                    v.GetString("network"),
		Manager:                    v.GetString("manager"),
		Nodes:                      v.GetInt("nodes"),
		Threshold:                  v.GetInt("threshold"),
		SharedDir:                  v.GetString("shared-dir"),
		ArtifactsDir:               v.GetString("artifacts-dir"),
		LogLevel:                   v.GetString("log-level"),
		WaitReadiness:              v.GetDuration("wait-readiness"),
		WaitContrib:                v.GetDuration("wait-contrib"),
		WaitDecrypt:                v.GetDuration("wait-decrypt"),
		ReadinessDeadlineBlocks:    v.GetInt("readiness-deadline-blocks"),
		ContributionDeadlineBlocks: v.GetInt("contribution-deadline-blocks"),
		DisclosureAllowed:          v.GetBool("disclosure-allowed"),
	}
	if c.PrivKey == "" {
		return nil, fmt.Errorf("--privkey required")
	}
	if c.Network != "" {
		if _, err := config.NetworkByName(c.Network); err != nil {
			return nil, err
		}
	}
	if c.Manager == "" && c.Network == "" {
		return nil, fmt.Errorf("--manager or --network required")
	}
	if c.Nodes < 1 || c.Threshold < 1 || c.Threshold > c.Nodes {
		return nil, fmt.Errorf("invalid --nodes / --threshold")
	}
	return c, nil
}

func main() {
	c, err := loadCfg()
	if err != nil {
		fmt.Fprintf(os.Stderr, "dkg-runner: %v\n", err)
		os.Exit(1)
	}
	log.Init(c.LogLevel, "stdout", nil)
	if c.ArtifactsDir != "" {
		circuits.BaseDir = c.ArtifactsDir
	} else if d := os.Getenv("DAVINCI_DKG_ARTIFACTS_DIR"); d != "" {
		circuits.BaseDir = d
	}
	if err := runScenario(c); err != nil {
		log.Errorw(err, "scenario failed", "err_msg", err.Error())
		fmt.Fprintf(os.Stderr, "FATAL ERROR: %v\n", err)
		os.Exit(1)
	}
	log.Infow("✓  DKG scenario completed successfully")
}

// runScenario executes the complete E2E DKG + threshold decryption scenario.
func runScenario(c *cfg) error {
	ctx := context.Background()

	// ── connect ──────────────────────────────────────────────────────────────
	// Resolve manager address: explicit --manager flag takes precedence; when
	// absent the network preset is used. All verifier addresses are derived
	// from the manager's public immutable fields by web3.New().
	managerAddr := c.Manager
	if managerAddr == "" && c.Network != "" {
		dep, err := config.NetworkByName(c.Network)
		if err != nil {
			return fmt.Errorf("resolve network: %w", err)
		}
		managerAddr = dep.Manager.Hex()
	}
	addrs := nodetypes.ContractAddresses{
		Manager: common.HexToAddress(managerAddr),
	}
	contracts, err := web3.New([]string{c.RPC}, addrs)
	if err != nil {
		return fmt.Errorf("web3: %w", err)
	}
	txm, err := txmanager.New(contracts.Pool().Current, contracts.ChainID, c.PrivKey)
	if err != nil {
		return fmt.Errorf("txmanager: %w", err)
	}
	manager, err := gtypes.NewDKGManager(addrs.Manager, contracts.Client())
	if err != nil {
		return err
	}
	callOpts := &bind.CallOpts{Context: ctx}
	n, t := uint16(c.Nodes), uint16(c.Threshold)
	log.Infow("runner connected", "organiser", txm.Address(), "nodes", n, "threshold", t)

	// ── 1. create round ──────────────────────────────────────────────────────
	head, err := contracts.Client().BlockNumber(ctx)
	if err != nil {
		return err
	}
	prefix, err := manager.ROUNDPREFIX(callOpts)
	if err != nil {
		return err
	}
	nonce0, err := manager.RoundNonce(callOpts)
	if err != nil {
		return err
	}

	auth, err := txm.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	// Lottery: each active node passes a keccak256(seed ‖ addr) < threshold
	// check independently. We oversubscribe by 1.5× so spare eligibles can
	// absorb last-moment no-shows without triggering an extendRegistration
	// reseed. On average ~1.5 × committeeSize nodes pass the check; the first
	// committeeSize to call claimSlot fill the committee.
	const lotteryAlphaBps uint16 = 15000 // α = 1.5
	const seedDelay uint16 = 1           // 1 block gap between createRound and seedBlock
	tx, err := manager.CreateRound(auth, t, n, t,
		lotteryAlphaBps, seedDelay,
		head+uint64(c.ReadinessDeadlineBlocks),
		head+uint64(c.ContributionDeadlineBlocks),
		c.DisclosureAllowed)
	if err != nil {
		return fmt.Errorf("create round: %w", err)
	}
	if err := txm.WaitTxByHash(tx.Hash(), 60*time.Second); err != nil {
		return err
	}
	roundID := makeRoundID(prefix, nonce0+1)
	log.Infow("round created", "id", roundHex(roundID))

	// ── 2. wait for committee to self-fill via claimSlot ─────────────────────
	log.Infow("waiting for committee to fill...", "want", n)
	if err := poll(c.WaitReadiness, "claimSlot", func() (bool, error) {
		r, err := contracts.GetRound(ctx, roundID)
		if err != nil {
			return false, err
		}
		log.Debugw("claims", "count", r.ClaimedCount, "want", n)
		return r.ClaimedCount >= n, nil
	}); err != nil {
		return err
	}
	committee, err := contracts.SelectedParticipants(ctx, roundID)
	if err != nil {
		return err
	}
	log.Infow("committee filled", "committee", addrList(committee))

	// ── 4. wait for t contributions, then until count is stable ─────────────
	log.Infow("waiting for contributions...", "want", t)
	var lastCount uint16
	var stableTicks int
	if err := poll(c.WaitContrib, "contributions", func() (bool, error) {
		r, err := contracts.GetRound(ctx, roundID)
		if err != nil {
			return false, err
		}
		log.Debugw("contributions", "count", r.ContributionCount, "want", t, "stable", stableTicks)
		if r.ContributionCount < t {
			lastCount = r.ContributionCount
			stableTicks = 0
			return false, nil
		}
		if r.ContributionCount == lastCount {
			stableTicks++
		} else {
			stableTicks = 0
			lastCount = r.ContributionCount
		}
		return stableTicks >= 2, nil
	}); err != nil {
		return err
	}

	// ── 5. finalize round ─────────────────────────────────────────────────────
	log.Infow("finalizing round...")
	shareCommitments, err := finalizeRound(ctx, contracts, manager, txm, roundID, t, n, committee)
	if err != nil {
		return fmt.Errorf("finalize: %w", err)
	}
	log.Infow("round finalized")

	// ── 6. encrypt test message ───────────────────────────────────────────────
	pk, err := collectivePK(ctx, contracts, manager, roundID, t, n, committee)
	if err != nil {
		return fmt.Errorf("recover PK: %w", err)
	}
	_ = shareCommitments

	l := group.ScalarField()
	m, err := rand.Int(rand.Reader, l)
	if err != nil {
		return err
	}
	if m.Sign() == 0 {
		m.SetInt64(42)
	}
	r2, err := rand.Int(rand.Reader, l)
	if err != nil {
		return err
	}
	if r2.Sign() == 0 {
		r2.SetInt64(1)
	}

	c1pt := group.Generator()
	c1pt.ScalarBaseMult(r2)
	pkNative, err := group.Decode(pk)
	if err != nil {
		return fmt.Errorf("decode PK: %w", err)
	}
	rPK := group.NewPoint()
	rPK.ScalarMult(pkNative, r2)
	mG := group.Generator()
	mG.ScalarBaseMult(m)
	c2pt := group.NewPoint()
	c2pt.Set(mG)
	c2pt.Add(c2pt, rPK)

	c1enc := group.Encode(c1pt)
	c2enc := group.Encode(c2pt)
	log.Infow("ciphertext created", "m", m.String(), "C1x", c1enc.X.String())

	if err := os.MkdirAll(c.SharedDir, 0o755); err != nil {
		return err
	}
	ctFile := CiphertextFile{CiphertextIndex: 1, C1X: c1enc.X.String(), C1Y: c1enc.Y.String()}
	ctData, _ := json.MarshalIndent(ctFile, "", "  ")
	ctPath := filepath.Join(c.SharedDir, fmt.Sprintf("ciphertext-%x.json", roundID))
	if err := os.WriteFile(ctPath, ctData, 0o644); err != nil {
		return err
	}
	log.Infow("ciphertext published to shared dir", "path", ctPath)

	// ── 7. wait for t partial decryptions ─────────────────────────────────────
	log.Infow("waiting for partial decryptions...", "want", t)
	if err := poll(c.WaitDecrypt, "partial decryptions", func() (bool, error) {
		count := uint16(0)
		for _, addr := range committee {
			rec, err := manager.GetPartialDecryption(&bind.CallOpts{Context: ctx}, roundID, addr, 1)
			if err == nil && rec.Accepted {
				count++
			}
		}
		log.Debugw("decryptions", "count", count, "want", t)
		return count >= t, nil
	}); err != nil {
		return err
	}

	// ── 8. combine decryptions ────────────────────────────────────────────────
	log.Infow("combining partial decryptions...")
	if err := combineDecryptions(ctx, contracts, manager, txm, roundID, t, committee, c1enc, c2enc, m); err != nil {
		return fmt.Errorf("combine: %w", err)
	}

	// ── 9. verify ─────────────────────────────────────────────────────────────
	combined, err := contracts.GetCombinedDecryption(ctx, roundID, 1)
	if err != nil {
		return fmt.Errorf("get combined: %w", err)
	}
	if !combined.Completed {
		return fmt.Errorf("combined decryption not marked completed")
	}
	// The combine proof was already verified on-chain (the tx would have reverted
	// otherwise) and the proof's public inputs include the plaintext scalar `m`,
	// so reaching this point means the on-chain decryption matched the original.
	log.Infow("✓  plaintext verified", "m", m.String())
	return nil
}

// ── finalize ─────────────────────────────────────────────────────────────────

// finalizeRound builds and submits the finalize proof from on-chain commitment
// points (public data fetched from contribution calldata).
func finalizeRound(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	txm *txmanager.Manager,
	roundID [12]byte,
	t, n uint16,
	committee []common.Address,
) ([]nodetypes.CurvePoint, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	acceptedIdxs := make([]uint16, 0, n)
	allPoints := make([][]nodetypes.CurvePoint, 0, n)

	for i, addr := range committee {
		rec, err := m.GetContribution(callOpts, roundID, addr)
		if err != nil || !rec.Accepted {
			continue
		}
		pts, err := commitmentPointsFromCalldata(ctx, c, addr, t)
		if err != nil {
			return nil, fmt.Errorf("commitment points for %s: %w", addr, err)
		}
		acceptedIdxs = append(acceptedIdxs, uint16(i+1))
		allPoints = append(allPoints, pts)
	}
	if len(acceptedIdxs) < int(t) {
		return nil, fmt.Errorf("only %d/%d accepted contributions", len(acceptedIdxs), t)
	}

	roundHash := roundScalar(roundID)
	asgn := finalize.CommitmentPointsAssignment{
		RoundHash:          roundHash,
		Threshold:          t,
		CommitteeSize:      n,
		ParticipantIndexes: acceptedIdxs,
		ContributionPoints: allPoints,
	}
	witness, pi, err := finalize.BuildWitnessFromCommitmentPoints(asgn)
	if err != nil {
		return nil, fmt.Errorf("build finalize witness: %w", err)
	}
	runtime, err := finalize.Artifacts.LoadOrSetupForCircuit(ctx, &finalize.FinalizeCircuit{})
	if err != nil {
		return nil, fmt.Errorf("load finalize circuit: %w", err)
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return nil, fmt.Errorf("prove finalize: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return nil, err
	}
	inputBytes, err := encodePublicWitness(pi.PublicWitness())
	if err != nil {
		return nil, err
	}
	transcriptBytes, err := encodeWords(pi.TranscriptScalars()...)
	if err != nil {
		return nil, err
	}

	auth, err := txm.NewTransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := m.FinalizeRound(auth, roundID,
		common.BigToHash(pi.AggregateHash),
		common.BigToHash(pi.CollectivePublicKey),
		common.BigToHash(pi.ShareCommitmentHash),
		transcriptBytes, proofBytes, inputBytes,
	)
	if err != nil {
		return nil, err
	}
	if err := txm.WaitTxByHash(tx.Hash(), 120*time.Second); err != nil {
		return nil, err
	}
	return pi.ShareCommitments, nil
}

// commitmentPointsFromCalldata scans recent blocks for the submitContribution
// transaction from `contributor` and returns its t Feldman commitment points.
func commitmentPointsFromCalldata(
	ctx context.Context,
	c *web3.Contracts,
	contributor common.Address,
	t uint16,
) ([]nodetypes.CurvePoint, error) {
	client := c.Client()
	chainID := new(big.Int).SetUint64(c.ChainID)
	latest, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	start := uint64(0)
	if latest > 2000 {
		start = latest - 2000
	}
	signer := ethtypes.NewCancunSigner(chainID)
	for blk := start; blk <= latest; blk++ {
		block, err := client.BlockByNumber(ctx, new(big.Int).SetUint64(blk))
		if err != nil {
			continue
		}
		for _, tx := range block.Transactions() {
			if tx.To() == nil || *tx.To() != c.Addresses.Manager {
				continue
			}
			from, err := ethtypes.Sender(signer, tx)
			if err != nil || from != contributor {
				continue
			}
			pts, err := parseCommitmentPoints(tx.Data(), t)
			if err != nil {
				continue
			}
			return pts, nil
		}
	}
	return nil, fmt.Errorf("submitContribution tx not found for %s", contributor)
}

// parseCommitmentPoints extracts the first t Feldman commitment points from
// the submitContribution calldata transcript.
// submitContribution(roundId, index, commitmentsHash, encryptedSharesHash,
//
//	commitment0X, commitment0Y, transcript, proof, input)
//
// The transcript is the 7th param (index 6); its ABI offset word sits at
// payload bytes 192-224 (6 × 32).
func parseCommitmentPoints(data []byte, t uint16) ([]nodetypes.CurvePoint, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("data too short")
	}
	payload := data[4:]
	// transcript is the 7th parameter (index 6, after commitment0X and commitment0Y),
	// offset is at head bytes 192..223 (9 params × 32 = 288-byte head; word 6 = bytes 192-224).
	if len(payload) < 224 {
		return nil, fmt.Errorf("payload head too short")
	}
	tOffset := int64(new(big.Int).SetBytes(pad32(payload[192:224])).Uint64())
	if int(tOffset)+32 > len(payload) {
		return nil, fmt.Errorf("transcript offset out of bounds")
	}
	tLen := int64(new(big.Int).SetBytes(pad32(payload[tOffset : tOffset+32])).Uint64())
	tStart := tOffset + 32
	if tStart+tLen > int64(len(payload)) {
		return nil, fmt.Errorf("transcript out of bounds")
	}
	tr := payload[tStart : tStart+tLen]
	// commitmentPoints: first MaxCoefficients points = MaxCoefficients*64 bytes
	if len(tr) < contribution.MaxCoefficients*64 {
		return nil, fmt.Errorf("transcript too short")
	}
	pts := make([]nodetypes.CurvePoint, t)
	for k := uint16(0); k < t; k++ {
		x := new(big.Int).SetBytes(tr[k*64 : k*64+32])
		y := new(big.Int).SetBytes(tr[k*64+32 : k*64+64])
		pts[k] = nodetypes.CurvePoint{X: x, Y: y}
	}
	return pts, nil
}

// ── collective public key ─────────────────────────────────────────────────────

// collectivePK computes PK = Cbar(0) = Σ C_i(0) by reading commitment
// point 0 from each accepted contributor's calldata.
func collectivePK(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	roundID [12]byte,
	t, n uint16,
	committee []common.Address,
) (nodetypes.CurvePoint, error) {
	_ = n
	callOpts := &bind.CallOpts{Context: ctx}
	sum := group.NewPoint()
	sum.SetZero()

	for _, addr := range committee {
		rec, err := m.GetContribution(callOpts, roundID, addr)
		if err != nil || !rec.Accepted {
			continue
		}
		pts, err := commitmentPointsFromCalldata(ctx, c, addr, t)
		if err != nil {
			return nodetypes.CurvePoint{}, err
		}
		if len(pts) == 0 {
			continue
		}
		pt, err := group.Decode(pts[0])
		if err != nil {
			return nodetypes.CurvePoint{}, err
		}
		sum.Add(sum, pt)
	}
	return group.Encode(sum), nil
}

// ── combine decryptions ───────────────────────────────────────────────────────

func combineDecryptions(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	txm *txmanager.Manager,
	roundID [12]byte,
	t uint16,
	committee []common.Address,
	c1, c2 nodetypes.CurvePoint,
	plaintext *big.Int,
) error {
	callOpts := &bind.CallOpts{Context: ctx}
	idxs := make([]uint16, 0, t)
	deltas := make([]nodetypes.CurvePoint, 0, t)

	for i, addr := range committee {
		if uint16(len(idxs)) >= t {
			break
		}
		rec, err := m.GetPartialDecryption(callOpts, roundID, addr, 1)
		if err != nil || !rec.Accepted {
			continue
		}
		idxs = append(idxs, uint16(i+1))
		deltas = append(deltas, nodetypes.CurvePoint{X: rec.Delta.X, Y: rec.Delta.Y})
	}
	if uint16(len(idxs)) < t {
		return fmt.Errorf("only %d/%d partial decryptions", len(idxs), t)
	}

	// Lagrange-interpolate Δ = sk*C1 from deltas (for verification inside witness builder).
	indexes := ccommon.Uint16sToBigInts(idxs)
	_, err := ccommon.InterpolatePointsAtZeroNative(indexes, deltas)
	if err != nil {
		return fmt.Errorf("interpolate check: %w", err)
	}

	asgn := decryptcombine.Assignment{
		RoundHash:          roundScalar(roundID),
		Threshold:          t,
		CiphertextC1:       c1,
		CiphertextC2:       c2,
		ParticipantIndexes: idxs,
		PartialDecryptions: deltas,
		Plaintext:          plaintext,
	}
	witness, pi, err := decryptcombine.BuildWitness(asgn)
	if err != nil {
		return fmt.Errorf("build combine witness: %w", err)
	}
	runtime, err := decryptcombine.Artifacts.LoadOrSetupForCircuit(ctx, &decryptcombine.DecryptCombineCircuit{})
	if err != nil {
		return err
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return fmt.Errorf("prove combine: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return err
	}
	inputBytes, err := encodePublicWitness(pi.PublicWitness())
	if err != nil {
		return err
	}
	transcriptBytes, err := encodeWords(pi.TranscriptScalars()...)
	if err != nil {
		return err
	}
	auth, err := txm.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := m.CombineDecryption(auth, roundID, 1,
		common.BigToHash(pi.CombineHash),
		common.BigToHash(pi.PlaintextHash),
		transcriptBytes, proofBytes, inputBytes,
	)
	if err != nil {
		return err
	}
	return txm.WaitTxByHash(tx.Hash(), 120*time.Second)
}

// ── helpers ───────────────────────────────────────────────────────────────────

func poll(timeout time.Duration, label string, cond func() (bool, error)) error {
	deadline := time.Now().Add(timeout)
	for {
		if time.Now().After(deadline) {
			return fmt.Errorf("timeout waiting for %s", label)
		}
		ok, err := cond()
		if err != nil {
			log.Warnw("poll error", "label", label, "err", err)
		}
		if ok {
			return nil
		}
		time.Sleep(5 * time.Second)
	}
}

func makeRoundID(prefix uint32, nonce uint64) [12]byte {
	var id [12]byte
	binary.BigEndian.PutUint32(id[:4], prefix)
	binary.BigEndian.PutUint64(id[4:], nonce)
	return id
}

func roundHex(id [12]byte) string      { return fmt.Sprintf("%x", id) }
func roundScalar(id [12]byte) *big.Int { return new(big.Int).SetBytes(id[:]) }

func addrList(addrs []common.Address) []string {
	s := make([]string, len(addrs))
	for i, a := range addrs {
		s[i] = a.Hex()
	}
	return s
}

func pad32(b []byte) []byte {
	if len(b) >= 32 {
		return b[len(b)-32:]
	}
	p := make([]byte, 32)
	copy(p[32-len(b):], b)
	return p
}
