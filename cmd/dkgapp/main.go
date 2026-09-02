// dkgapp is the application-side companion of davinci-dkg-node: it registers
// applications against a live epoch (public derivation or organizer
// co-decryption), encrypts and submits ciphertexts, posts organizer shares
// and reads back combined plaintexts. It is what an integrator or an
// election organizer runs; committee members run davinci-dkg-node.
package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"math"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/vocdoni/davinci-dkg/circuits/partialdecrypt"
	"github.com/vocdoni/davinci-dkg/config"
	"github.com/vocdoni/davinci-dkg/crypto/elgamal"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/crypto/schnorr"
	"github.com/vocdoni/davinci-dkg/internal/protocol"
	"github.com/vocdoni/davinci-dkg/log"
	"github.com/vocdoni/davinci-dkg/prover"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
)

const usage = `usage: dkgapp [-rpc url[,url]] [-network name | -manager 0x..] [-privkey hex] <command> [flags]

commands:
  epoch      [-epoch id]                                   print an epoch (default: latest)
  register    -epoch id -aid hex32 [-codec] [-org-secret hex] [-submitter 0x..] [-max n]
                                                            register an application (mode 0, or mode 1 with -codec)
  encrypt     -epoch id -aid hex32 -m int [-org-secret hex]
                                                            encrypt m under PK_aid with a proof of knowledge of
                                                            its randomness, submit it (the chain assigns the
                                                            index) and, for mode 1 with -org-secret, post the
                                                            organizer share right away (omit it to withhold)
  share       -epoch id -aid hex32 -index n -org-secret hex
                                                            post the organizer share of a mode-1 ciphertext
                                                            that was submitted earlier (releases decryption)
  plaintext   -epoch id -aid hex32 -index n [-wait dur]    read (or wait for) the combined plaintext

Proofs are generated with the pinned release circuit artifacts; point
DAVINCI_ARTIFACTS_DIR at a directory holding them (see README, "Circuits").

Every flag has a DAVINCI_DKG_* environment equivalent for the global options
(DAVINCI_DKG_WEB3_RPC, DAVINCI_DKG_NETWORK, DAVINCI_DKG_MANAGER, DAVINCI_DKG_PRIVKEY).
`

type app struct {
	ctx        context.Context
	contracts  *web3.Contracts
	manager    *gtypes.DKGManager
	appManager *gtypes.DKGAppManager
	txm        *txmanager.Manager
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	global := flag.NewFlagSet("dkgapp", flag.ContinueOnError)
	rpc := global.String("rpc", envOr("DAVINCI_DKG_WEB3_RPC", ""), "comma-separated JSON-RPC endpoints")
	network := global.String("network", envOr("DAVINCI_DKG_NETWORK", ""), "well-known network preset (e.g. sepolia)")
	managerAddr := global.String("manager", envOr("DAVINCI_DKG_MANAGER", ""), "DKGManager address (overrides -network)")
	privkey := global.String("privkey", envOr("DAVINCI_DKG_PRIVKEY", ""), "hex private key used to sign transactions")
	global.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	if err := global.Parse(os.Args[1:]); err != nil {
		return err
	}
	if global.NArg() == 0 {
		global.Usage()
		return fmt.Errorf("missing command")
	}
	log.Init("info", "stderr", nil)

	if *managerAddr == "" && *network != "" {
		dep, err := config.NetworkByName(*network)
		if err != nil {
			return err
		}
		*managerAddr = dep.Manager.Hex()
	}
	if *managerAddr == "" || *rpc == "" {
		return fmt.Errorf("-rpc and -manager (or -network) are required")
	}
	contracts, err := web3.New(strings.Split(*rpc, ","), types.ContractAddresses{Manager: common.HexToAddress(*managerAddr)})
	if err != nil {
		return err
	}
	manager, err := gtypes.NewDKGManager(contracts.Addresses.Manager, contracts.PooledBackend())
	if err != nil {
		return err
	}
	appManager, err := gtypes.NewDKGAppManager(contracts.Addresses.AppManager, contracts.PooledBackend())
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	a := &app{ctx: ctx, contracts: contracts, manager: manager, appManager: appManager}
	if *privkey != "" {
		a.txm, err = txmanager.New(contracts.Pool().Current, contracts.ChainID, *privkey)
		if err != nil {
			return err
		}
		a.txm.Start(ctx)
		defer a.txm.Stop()
	}

	cmd, args := global.Arg(0), global.Args()[1:]
	switch cmd {
	case "epoch":
		return a.cmdEpoch(args)
	case "register":
		return a.cmdRegister(args)
	case "encrypt":
		return a.cmdEncrypt(args)
	case "share":
		return a.cmdShare(args)
	case "plaintext":
		return a.cmdPlaintext(args)
	default:
		global.Usage()
		return fmt.Errorf("unknown command %q", cmd)
	}
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// ── epoch ──────────────────────────────────────────────────────────────────

func (a *app) cmdEpoch(args []string) error {
	fs := flag.NewFlagSet("epoch", flag.ContinueOnError)
	epochFlag := fs.String("epoch", "latest", "12-byte epoch id (hex) or 'latest'")
	if err := fs.Parse(args); err != nil {
		return err
	}
	id, err := a.epochID(*epochFlag)
	if err != nil {
		return err
	}
	e, err := a.contracts.GetEpoch(a.ctx, id)
	if err != nil {
		return err
	}
	pk, err := a.manager.GetCollectivePublicKey(a.callOpts(), id)
	if err != nil {
		return err
	}
	phases := []string{"None", "CommitteeSelection", "KeyAssembly", "Live", "Aborted", "Completed"}
	phase := fmt.Sprint(e.Status)
	if int(e.Status) < len(phases) {
		phase = phases[e.Status]
	}
	fmt.Printf("epoch      %x\nphase      %s\nthreshold  %d/%d (min %d)\nclaimed    %d\ncontribs   %d\nciphertexts %d\nstartBlock %d\nPK_ep      (%s, %s)\n",
		id, phase, e.Policy.Threshold, e.Policy.CommitteeSize, e.Policy.MinValidContributions,
		e.ClaimedCount, e.ContributionCount, e.CiphertextCount, e.StartBlock, pk.X, pk.Y)
	return nil
}

// ── register ───────────────────────────────────────────────────────────────

func (a *app) cmdRegister(args []string) error {
	fs := flag.NewFlagSet("register", flag.ContinueOnError)
	epochFlag := fs.String("epoch", "latest", "epoch id (hex) or 'latest'")
	aidFlag := fs.String("aid", "", "32-byte application id (hex), must be non-zero and below the BN254 scalar field")
	codec := fs.Bool("codec", false, "organizer co-decryption mode (mode 1)")
	orgSecret := fs.String("org-secret", "", "organizer secret scalar (hex); generated when omitted")
	submitter := fs.String("submitter", "", "only this address may submit ciphertexts (default: anyone)")
	maxCt := fs.Uint("max", 0, "maximum ciphertexts (0 = unlimited)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if a.txm == nil {
		return fmt.Errorf("-privkey is required")
	}
	id, err := a.epochID(*epochFlag)
	if err != nil {
		return err
	}
	aid, err := parseAid(*aidFlag)
	if err != nil {
		return err
	}
	policy := gtypes.DKGTypesAppPolicy{MaxCiphertexts: uint16(*maxCt)}
	if *submitter != "" {
		policy.AuthorizedSubmitter = common.HexToAddress(*submitter)
	}
	auth, err := a.txm.NewTransactOpts(a.ctx)
	if err != nil {
		return err
	}
	var tx *ethtypes.Transaction
	if !*codec {
		tx, err = a.appManager.RegisterApplication(auth, id, aid, policy)
	} else {
		var sk *big.Int
		var generated bool
		sk, generated, err = organizerSecret(*orgSecret)
		if err != nil {
			return err
		}
		var pkX, pkY *big.Int
		var proof schnorr.OrganizerProof
		pkX, pkY, proof, err = schnorr.ProveOrganizerRegister(sk, id, aid)
		if err != nil {
			return err
		}
		if generated {
			fmt.Printf("organizer secret (keep it, needed to decrypt): %x\n", sk)
		}
		fmt.Printf("organizer key  (%s, %s)\n", pkX, pkY)
		tx, err = a.appManager.RegisterApplicationCoDec(auth, id, aid, policy, pkX, pkY, proof.Ax, proof.Ay, proof.Z)
	}
	if err != nil {
		return fmt.Errorf("register: %w", err)
	}
	if err := a.wait(tx); err != nil {
		return err
	}
	fmt.Printf("application %x registered in epoch %x (tx %s)\n", aid, id, tx.Hash().Hex())
	return nil
}

// ── encrypt ────────────────────────────────────────────────────────────────

func (a *app) cmdEncrypt(args []string) error {
	fs := flag.NewFlagSet("encrypt", flag.ContinueOnError)
	epochFlag := fs.String("epoch", "latest", "epoch id (hex) or 'latest'")
	aidFlag := fs.String("aid", "", "application id (hex); all-zero for the epoch key itself")
	mFlag := fs.String("m", "", "plaintext as a decimal integer (< 2^50 for committee recovery)")
	orgSecret := fs.String("org-secret", "", "organizer secret (hex); required for mode-1 applications")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if a.txm == nil {
		return fmt.Errorf("-privkey is required")
	}
	m, ok := new(big.Int).SetString(*mFlag, 10)
	if !ok || m.Sign() < 0 {
		return fmt.Errorf("-m must be a non-negative decimal integer")
	}
	id, err := a.epochID(*epochFlag)
	if err != nil {
		return err
	}
	aid, err := parseAidAllowZero(*aidFlag)
	if err != nil {
		return err
	}
	pkEpRaw, err := a.manager.GetCollectivePublicKey(a.callOpts(), id)
	if err != nil {
		return err
	}
	pkEp := types.CurvePoint{X: pkEpRaw.X, Y: pkEpRaw.Y}
	var mode uint8
	pkAid := pkEp
	if aid != ([32]byte{}) {
		rec, err := a.appManager.GetApplication(a.callOpts(), id, aid)
		if err != nil {
			return err
		}
		if !rec.Exists {
			return fmt.Errorf("application %x is not registered in epoch %x", aid, id)
		}
		mode = rec.Mode
		pkAid, err = elgamal.ApplicationKey(pkEp, mode, rec.DerivationS, types.CurvePoint{X: rec.OrganizerPK.X, Y: rec.OrganizerPK.Y})
		if err != nil {
			return err
		}
	}
	c1, c2, pok, err := elgamal.EncryptWithProof(id, aid, pkAid, m)
	if err != nil {
		return err
	}
	auth, err := a.txm.NewTransactOpts(a.ctx)
	if err != nil {
		return err
	}
	tx, err := a.manager.SubmitCiphertext(auth, id, aid, c1.X, c1.Y, c2.X, c2.Y, pok.A.X, pok.A.Y, pok.Z)
	if err != nil {
		return fmt.Errorf("submit ciphertext: %w", err)
	}
	if err := a.wait(tx); err != nil {
		return err
	}
	index, err := a.ciphertextIndex(tx.Hash())
	if err != nil {
		return err
	}
	fmt.Printf("ciphertext %d submitted for %x (tx %s)\n", index, aid, tx.Hash().Hex())
	if mode != 1 {
		return nil
	}
	if *orgSecret == "" {
		fmt.Printf("organizer share withheld; release it with: share -epoch %x -aid %x -index %d -org-secret ...\n", id, aid, index)
		return nil
	}
	sk, _, err := organizerSecret(*orgSecret)
	if err != nil {
		return err
	}
	return a.submitOrganizerShare(id, aid, index, c1, c2, sk)
}

func (a *app) cmdShare(args []string) error {
	fs := flag.NewFlagSet("share", flag.ContinueOnError)
	epochFlag := fs.String("epoch", "latest", "epoch id (hex) or 'latest'")
	aidFlag := fs.String("aid", "", "application id (hex)")
	indexFlag := fs.Uint("index", 1, "ciphertext index assigned at submission")
	orgSecret := fs.String("org-secret", "", "organizer secret (hex)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if a.txm == nil {
		return fmt.Errorf("-privkey is required")
	}
	if *indexFlag == 0 || *indexFlag > math.MaxUint16 {
		return fmt.Errorf("-index must be in [1, %d]", math.MaxUint16)
	}
	id, err := a.epochID(*epochFlag)
	if err != nil {
		return err
	}
	aid, err := parseAid(*aidFlag)
	if err != nil {
		return err
	}
	sk, _, err := organizerSecret(*orgSecret)
	if err != nil {
		return err
	}
	ep, err := a.manager.GetEpoch(a.callOpts(), id)
	if err != nil {
		return err
	}
	index := uint16(*indexFlag)
	it, err := a.manager.FilterCiphertextSubmitted(
		&bind.FilterOpts{Context: a.ctx, Start: ep.StartBlock}, [][12]byte{id}, [][32]byte{aid}, []uint16{index},
	)
	if err != nil {
		return fmt.Errorf("scan ciphertext events: %w", err)
	}
	defer func() { _ = it.Close() }()
	if !it.Next() {
		return fmt.Errorf("ciphertext %d for application %x not found in epoch %x", index, aid, id)
	}
	ev := it.Event
	c1 := types.CurvePoint{X: ev.C1x, Y: ev.C1y}
	c2 := types.CurvePoint{X: ev.C2x, Y: ev.C2y}
	return a.submitOrganizerShare(id, aid, index, c1, c2, sk)
}

func (a *app) submitOrganizerShare(id [12]byte, aid [32]byte, index uint16, c1, c2 types.CurvePoint, sk *big.Int) error {
	nonce, err := rand.Int(rand.Reader, group.ScalarField())
	if err != nil {
		return err
	}
	witness, pi, err := partialdecrypt.BuildWitness(partialdecrypt.Assignment{
		RoundHash:        new(big.Int).SetBytes(id[:]),
		Aid:              new(big.Int).SetBytes(aid[:]),
		CtIdx:            new(big.Int).SetUint64(uint64(index)),
		Role:             big.NewInt(int64(protocol.RoleOrganizer)),
		ParticipantIndex: 0,
		Base:             c1,
		Secret:           sk,
		Nonce:            nonce,
	})
	if err != nil {
		return err
	}
	runtime, err := partialdecrypt.Artifacts.LoadPinned(a.ctx, &partialdecrypt.PartialDecryptCircuit{})
	if err != nil {
		return fmt.Errorf("load partial-decrypt circuit: %w", err)
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return fmt.Errorf("prove organizer share: %w", err)
	}
	proofBytes, err := prover.MarshalSolidityProof(proof)
	if err != nil {
		return err
	}
	inputBytes, err := prover.EncodePublicWitness(pi.PublicWitness())
	if err != nil {
		return err
	}
	auth, err := a.txm.NewTransactOpts(a.ctx)
	if err != nil {
		return err
	}
	tx, err := a.appManager.SubmitOrganizerShare(auth, id, aid, index, c1.X, c1.Y, c2.X, c2.Y, pi.Delta.X, pi.Delta.Y, proofBytes, inputBytes)
	if err != nil {
		return fmt.Errorf("submit organizer share: %w", err)
	}
	if err := a.wait(tx); err != nil {
		return err
	}
	fmt.Printf("organizer share submitted (tx %s)\n", tx.Hash().Hex())
	return nil
}

// ── plaintext ──────────────────────────────────────────────────────────────

func (a *app) cmdPlaintext(args []string) error {
	fs := flag.NewFlagSet("plaintext", flag.ContinueOnError)
	epochFlag := fs.String("epoch", "latest", "epoch id (hex) or 'latest'")
	aidFlag := fs.String("aid", "", "application id (hex); all-zero for the epoch key itself")
	index := fs.Uint("index", 1, "ciphertext index")
	wait := fs.Duration("wait", 0, "poll until combined for up to this long (0 = read once)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	id, err := a.epochID(*epochFlag)
	if err != nil {
		return err
	}
	aid, err := parseAidAllowZero(*aidFlag)
	if err != nil {
		return err
	}
	deadline := time.Now().Add(*wait)
	for {
		rec, err := a.manager.GetCombinedDecryption(a.callOpts(), id, aid, uint16(*index))
		if err != nil {
			return err
		}
		if rec.Completed {
			fmt.Printf("%s\n", rec.Plaintext)
			return nil
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("ciphertext %d of %x not combined yet", *index, aid)
		}
		time.Sleep(5 * time.Second)
	}
}

// ── helpers ────────────────────────────────────────────────────────────────

func (a *app) callOpts() *bind.CallOpts { return &bind.CallOpts{Context: a.ctx} }

// ciphertextIndex reads the on-chain-assigned index from the submit receipt.
func (a *app) ciphertextIndex(txHash common.Hash) (uint16, error) {
	receipt, err := a.contracts.Client().TransactionReceipt(a.ctx, txHash)
	if err != nil {
		return 0, fmt.Errorf("receipt: %w", err)
	}
	for _, lg := range receipt.Logs {
		if ev, err := a.manager.ParseCiphertextSubmitted(*lg); err == nil {
			return ev.CiphertextIndex, nil
		}
	}
	return 0, fmt.Errorf("CiphertextSubmitted event not found")
}

func (a *app) wait(tx *ethtypes.Transaction) error {
	a.txm.RecordPending(tx)
	return a.txm.WaitTxByHash(tx.Hash(), 3*time.Minute)
}

func (a *app) epochID(s string) ([12]byte, error) {
	var id [12]byte
	if s == "latest" {
		nonce, err := a.manager.EpochNonce(a.callOpts())
		if err != nil {
			return id, err
		}
		if nonce == 0 {
			return id, fmt.Errorf("no epoch has been created yet")
		}
		prefix, err := a.manager.EPOCHPREFIX(a.callOpts())
		if err != nil {
			return id, err
		}
		return web3.EpochID(prefix, nonce), nil
	}
	b, err := hex.DecodeString(strings.TrimPrefix(s, "0x"))
	if err != nil || len(b) != 12 {
		return id, fmt.Errorf("epoch id must be 12 bytes of hex")
	}
	copy(id[:], b)
	return id, nil
}

func parseAidAllowZero(s string) ([32]byte, error) {
	var aid [32]byte
	if s == "" || s == "0" || s == "0x0" {
		return aid, nil
	}
	return parseAid(s)
}

// parseAid accepts a 32-byte hex id and enforces the on-chain rule that it is
// non-zero and below the BN254 scalar field (it is a proof public input).
func parseAid(s string) ([32]byte, error) {
	var aid [32]byte
	b, err := hex.DecodeString(strings.TrimPrefix(s, "0x"))
	if err != nil || len(b) == 0 || len(b) > 32 {
		return aid, fmt.Errorf("aid must be 1..32 bytes of hex")
	}
	copy(aid[32-len(b):], b)
	v := new(big.Int).SetBytes(aid[:])
	if v.Sign() == 0 || v.Cmp(group.BaseField()) >= 0 {
		return aid, fmt.Errorf("aid must be non-zero and below the BN254 scalar field (clear its top three bits)")
	}
	return aid, nil
}

func organizerSecret(hexSecret string) (sk *big.Int, generated bool, err error) {
	if hexSecret != "" {
		sk, ok := new(big.Int).SetString(strings.TrimPrefix(hexSecret, "0x"), 16)
		if !ok || sk.Sign() == 0 || sk.Cmp(group.ScalarField()) >= 0 {
			return nil, false, fmt.Errorf("org-secret must be a non-zero hex scalar below the BabyJubJub order")
		}
		return sk, false, nil
	}
	sk, err = rand.Int(rand.Reader, group.ScalarField())
	if err != nil {
		return nil, false, err
	}
	if sk.Sign() == 0 {
		sk.SetInt64(1)
	}
	return sk, true, nil
}
