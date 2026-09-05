// dkgapp is the application-side companion of davinci-dkg-node: it registers
// applications against a Live epoch's pool of committee keys, encrypts and
// submits ciphertexts, reveals the organizer secret that releases decryption
// and reads back combined plaintexts. It is what an integrator or an election
// organizer runs; committee members run davinci-dkg-node.
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
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/config"
	"github.com/vocdoni/davinci-dkg/crypto/elgamal"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/crypto/schnorr"
	"github.com/vocdoni/davinci-dkg/log"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
)

const usage = `usage: dkgapp [-rpc url[,url]] [-network name | -manager 0x..] [-privkey hex] <command> [flags]

commands:
  epoch      [-epoch id]                                   print an epoch and its key pool (default: latest)
  register   [-epoch id] -aid hex32 [-mode locked|automatic] [-org-secret hex]
              [-submitters 0x..,0x..] [-open] [-max n]
              [-decrypt-from rfc3339|duration] [-decrypt-until rfc3339|duration]
                                                            claim one of the epoch's committee keys for the
                                                            application; without -epoch the newest Live epoch
                                                            with an activated unclaimed key is used. In locked
                                                            mode the organizer key PK_org = sk_org*G is
                                                            published with a Schnorr proof of possession and
                                                            the secret is generated and printed when
                                                            -org-secret is omitted; in automatic mode there is
                                                            no organizer key at all
  encrypt     -epoch id -aid hex32 -m int                   encrypt m under PK_aid = P_j (+ PK_org when the
                                                            application is organizer-locked) and submit it;
                                                            the chain assigns the index
  reveal      -epoch id -aid hex32 -org-secret hex          publish sk_org once and for all: from then on the
                                                            committee decrypts every ciphertext of the
                                                            application by itself. Permissionless, one-shot,
                                                            no SNARK and no circuit artifacts
  plaintext   -epoch id -aid hex32 -index n [-wait dur]    read (or wait for) the combined plaintext

Losing sk_org before revealing it makes every ciphertext of an
organizer-locked application permanently undecryptable: the committee alone
only recovers its share of P_j.

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
	case "reveal":
		return a.cmdReveal(args)
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
	status, err := a.manager.GetPoolStatus(a.callOpts(), id)
	if err != nil {
		return err
	}
	phases := []string{"None", "CommitteeSelection", "KeyAssembly", "Live", "Aborted", "Completed"}
	phase := fmt.Sprint(e.Status)
	if int(e.Status) < len(phases) {
		phase = phases[e.Status]
	}
	fmt.Printf("epoch      %x\nphase      %s\nthreshold  %d/%d (min %d)\nclaimed    %d\ncontribs   %d\n"+
		"ciphertexts %d\nstartBlock %d\npool       %d/%d claimed, activated bitmap %08b\n",
		id, phase, e.Policy.Threshold, e.Policy.CommitteeSize, e.Policy.MinValidContributions,
		e.ClaimedCount, e.ContributionCount, e.CiphertextCount, e.StartBlock,
		status.NextIndex, ccommon.MaxK, status.Activated)
	for key := range uint8(ccommon.MaxK) {
		if status.Activated&(1<<key) == 0 {
			continue
		}
		x, y, keyErr := a.manager.GetPoolKey(a.callOpts(), id, key)
		if keyErr != nil {
			return keyErr
		}
		state := "claimed"
		if key >= status.NextIndex {
			state = "free"
		}
		fmt.Printf("  P_%d      (%s, %s) %s\n", key, x, y, state)
	}
	return nil
}

// ── register ───────────────────────────────────────────────────────────────

func (a *app) cmdRegister(args []string) error {
	fs := flag.NewFlagSet("register", flag.ContinueOnError)
	epochFlag := fs.String("epoch", "",
		"epoch id (hex); default: the newest Live epoch with an activated unclaimed pool key")
	aidFlag := fs.String("aid", "", "32-byte application id (hex), must be non-zero and below the BN254 scalar field")
	modeFlag := fs.String("mode", "locked", "'locked': PK_aid = P_j + PK_org and you reveal sk_org when decryption may start; "+
		"'automatic': no organizer key, the committee decrypts as soon as the ciphertexts land")
	orgSecret := fs.String("org-secret", "", "organizer secret scalar (hex); generated and printed when omitted (locked mode)")
	submitters := fs.String("submitters", "",
		"comma-separated addresses allowed to submit ciphertexts (default: the registering address only)")
	open := fs.Bool("open", false, "let anyone submit ciphertexts")
	maxCt := fs.Uint("max", 0, "maximum ciphertexts (0 = unlimited)")
	from := fs.String("decrypt-from", "",
		"earliest decryption time: RFC3339 timestamp or a duration from now such as 24h (default: immediately)")
	until := fs.String("decrypt-until", "",
		"decryption deadline: RFC3339 timestamp or a duration from now such as 48h (default: none)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if a.txm == nil {
		return fmt.Errorf("-privkey is required")
	}
	if *maxCt > math.MaxUint16 {
		return fmt.Errorf("-max must fit in a uint16")
	}
	policy := gtypes.DKGTypesAppPolicy{MaxCiphertexts: uint16(*maxCt), OpenSubmission: *open}
	switch *modeFlag {
	case "locked":
	case "automatic":
		policy.Mode = uint8(types.AppModeAutomatic)
	default:
		return fmt.Errorf("-mode must be 'locked' or 'automatic'")
	}
	for _, s := range strings.Split(*submitters, ",") {
		if s = strings.TrimSpace(s); s == "" {
			continue
		}
		if !common.IsHexAddress(s) {
			return fmt.Errorf("-submitters: %q is not an address", s)
		}
		policy.Submitters = append(policy.Submitters, common.HexToAddress(s))
	}
	now := time.Now()
	if *from != "" {
		opensAt, err := parseDeadline(*from, now)
		if err != nil {
			return fmt.Errorf("-decrypt-from: %w", err)
		}
		policy.DecryptNotBefore = opensAt
	}
	if *until != "" {
		deadline, err := parseDeadline(*until, now)
		if err != nil {
			return fmt.Errorf("-decrypt-until: %w", err)
		}
		policy.DecryptNotAfter = deadline
	}
	if policy.DecryptNotBefore != 0 && policy.DecryptNotAfter != 0 &&
		policy.DecryptNotAfter <= policy.DecryptNotBefore {
		return fmt.Errorf("-decrypt-until must be after -decrypt-from")
	}

	id, err := a.registrationEpoch(*epochFlag)
	if err != nil {
		return err
	}
	aid, err := parseAid(*aidFlag)
	if err != nil {
		return err
	}

	pkX, pkY := new(big.Int), new(big.Int)
	ax, ay, z := new(big.Int), new(big.Int), new(big.Int)
	if policy.Mode != uint8(types.AppModeAutomatic) {
		sk, generated, secretErr := organizerSecret(*orgSecret)
		if secretErr != nil {
			return secretErr
		}
		organizerX, organizerY, proof, proofErr := schnorr.ProveOrganizerRegister(sk, id, aid)
		if proofErr != nil {
			return proofErr
		}
		pkX, pkY = organizerX, organizerY
		ax, ay, z = proof.Ax, proof.Ay, proof.Z
		if generated {
			fmt.Printf("organizer secret: %x\n", sk)
			fmt.Println("WARNING: store this now. It is not derivable from anything on chain, and until")
			fmt.Println("         you publish it with \"reveal\" every ciphertext of this application")
			fmt.Println("         stays undecryptable — the committee alone only recovers its share of P_j.")
		}
		fmt.Printf("organizer key  (%s, %s)\n", pkX, pkY)
	}

	auth, err := a.txm.NewTransactOpts(a.ctx)
	if err != nil {
		return err
	}
	tx, err := a.appManager.RegisterApplication(auth, id, aid, policy, pkX, pkY, ax, ay, z)
	if err != nil {
		return fmt.Errorf("register: %w", err)
	}
	if err := a.wait(tx); err != nil {
		return err
	}
	rec, err := a.appManager.GetApplication(a.callOpts(), id, aid)
	if err != nil {
		return err
	}
	fmt.Printf("application %x registered in epoch %x on pool key %d (%s mode, tx %s)\n",
		aid, id, rec.PoolIndex, *modeFlag, tx.Hash().Hex())
	return nil
}

// registrationEpoch resolves the epoch a registration goes to: the one the
// caller named, or the newest Live epoch whose pool still has an activated,
// unclaimed key. Every other epoch would revert with PoolExhausted or
// PoolKeyNotActive.
func (a *app) registrationEpoch(explicit string) ([12]byte, error) {
	var id [12]byte
	if explicit != "" {
		return a.epochID(explicit)
	}
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
	// Only a handful of epochs are ever serviceable at once; the nodes keep
	// activating keys ahead in the newest ones.
	const lookback = 8
	for n := nonce; n > 0 && n+lookback > nonce; n-- {
		candidate := web3.EpochID(prefix, n)
		epoch, epochErr := a.manager.GetEpoch(a.callOpts(), candidate)
		if epochErr != nil || epoch.Status != uint8(types.EpochPhaseLive) {
			continue
		}
		status, statusErr := a.manager.GetPoolStatus(a.callOpts(), candidate)
		if statusErr != nil {
			continue
		}
		if status.NextIndex < ccommon.MaxK && status.Activated&(1<<status.NextIndex) != 0 {
			return candidate, nil
		}
	}
	return id, fmt.Errorf("no Live epoch has an activated unclaimed pool key; retry once the committee activates one")
}

// parseDeadline accepts a Go duration relative to now ("48h") or an RFC3339
// timestamp and returns unix seconds strictly in the future.
func parseDeadline(s string, now time.Time) (uint64, error) {
	var at time.Time
	if d, err := time.ParseDuration(s); err == nil {
		at = now.Add(d)
	} else if at, err = time.Parse(time.RFC3339, s); err != nil {
		return 0, fmt.Errorf("must be a duration such as 48h or an RFC3339 timestamp: %w", err)
	}
	if !at.After(now) {
		return 0, fmt.Errorf("must be in the future")
	}
	return uint64(at.Unix()), nil
}

// ── encrypt ────────────────────────────────────────────────────────────────

func (a *app) cmdEncrypt(args []string) error {
	fs := flag.NewFlagSet("encrypt", flag.ContinueOnError)
	epochFlag := fs.String("epoch", "latest", "epoch id (hex) or 'latest'")
	aidFlag := fs.String("aid", "", "application id (hex)")
	mFlag := fs.String("m", "", "plaintext as a decimal integer (< 2^50 for committee recovery)")
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
	aid, err := parseAid(*aidFlag)
	if err != nil {
		return err
	}
	pkAid, err := a.applicationKey(id, aid)
	if err != nil {
		return err
	}
	c1, c2, err := elgamal.Encrypt(pkAid, m)
	if err != nil {
		return err
	}
	auth, err := a.txm.NewTransactOpts(a.ctx)
	if err != nil {
		return err
	}
	tx, err := a.manager.SubmitCiphertext(auth, id, aid, c1.X, c1.Y, c2.X, c2.Y)
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
	return nil
}

// applicationKey is PK_aid = P_j (+ PK_org when the application is
// organizer-locked), with P_j the committee key the registration claimed.
func (a *app) applicationKey(id [12]byte, aid [32]byte) (types.CurvePoint, error) {
	rec, err := a.appManager.GetApplication(a.callOpts(), id, aid)
	if err != nil {
		return types.CurvePoint{}, err
	}
	if !rec.Exists {
		return types.CurvePoint{}, fmt.Errorf("application %x is not registered in epoch %x", aid, id)
	}
	x, y, err := a.manager.GetPoolKey(a.callOpts(), id, rec.PoolIndex)
	if err != nil {
		return types.CurvePoint{}, fmt.Errorf("pool key %d of epoch %x: %w", rec.PoolIndex, id, err)
	}
	return elgamal.ApplicationKey(
		types.CurvePoint{X: x, Y: y},
		types.CurvePoint{X: rec.OrganizerPK.X, Y: rec.OrganizerPK.Y},
	)
}

// ── reveal ─────────────────────────────────────────────────────────────────

// cmdReveal publishes sk_org for an organizer-locked application. It is the
// single organizer action of the whole protocol: permissionless, one-shot,
// and from then on the committee decrypts every ciphertext of the
// application — past and future — by itself. There is no SNARK here, so the
// organizer needs no circuit artifacts.
func (a *app) cmdReveal(args []string) error {
	fs := flag.NewFlagSet("reveal", flag.ContinueOnError)
	epochFlag := fs.String("epoch", "latest", "epoch id (hex) or 'latest'")
	aidFlag := fs.String("aid", "", "application id (hex)")
	orgSecret := fs.String("org-secret", "", "organizer secret (hex)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if a.txm == nil {
		return fmt.Errorf("-privkey is required")
	}
	if *orgSecret == "" {
		return fmt.Errorf("-org-secret is required")
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
	auth, err := a.txm.NewTransactOpts(a.ctx)
	if err != nil {
		return err
	}
	tx, err := a.appManager.RevealOrganizerSecret(auth, id, aid, sk)
	if err != nil {
		return fmt.Errorf("reveal organizer secret: %w", err)
	}
	if err := a.wait(tx); err != nil {
		return err
	}
	fmt.Printf("organizer secret revealed for %x; the committee can now decrypt (tx %s)\n", aid, tx.Hash().Hex())
	return nil
}

// ── plaintext ──────────────────────────────────────────────────────────────

func (a *app) cmdPlaintext(args []string) error {
	fs := flag.NewFlagSet("plaintext", flag.ContinueOnError)
	epochFlag := fs.String("epoch", "latest", "epoch id (hex) or 'latest'")
	aidFlag := fs.String("aid", "", "application id (hex)")
	index := fs.Uint("index", 1, "ciphertext index")
	wait := fs.Duration("wait", 0, "poll until combined for up to this long (0 = read once)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	id, err := a.epochID(*epochFlag)
	if err != nil {
		return err
	}
	aid, err := parseAid(*aidFlag)
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

// organizerSecret parses the caller-supplied organizer scalar, or draws a
// fresh one from crypto/rand. `generated` tells the caller to print it: it is
// the only copy, and losing it makes the application undecryptable.
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
