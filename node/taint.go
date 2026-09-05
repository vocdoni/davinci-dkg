package node

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/log"
)

// Tainted applications are remembered on disk so a restarted node does not
// repeat the 2^50 search that condemned them: a fleet restart would otherwise
// run one full-core search per node per poisoned application at once.

// taintKey names what an undecryptable ciphertext condemns for the rest of
// the epoch: always its (application, submitter) pair. Whoever produced it
// pays one 2^50 search per address, and the application's honest submitters
// keep their service. Entries without a submitter are a legacy on-disk form
// from before this rule (whole-application taints); they still load and
// match.
type taintKey struct {
	epoch     [12]byte
	aid       [32]byte
	submitter common.Address
}

func taintPath(datadir string) string {
	if datadir == "" {
		return ""
	}
	return filepath.Join(datadir, "tainted-apps.json")
}

// tainted reports whether the slot must be ignored: its application is
// tainted as a whole, or its submitter is tainted for that application.
func (n *Node) tainted(key ctKey, submitter common.Address) bool {
	if n.taints[taintKey{epoch: key.epoch, aid: key.aid}] {
		return true
	}
	return submitter != (common.Address{}) && n.taints[taintKey{epoch: key.epoch, aid: key.aid, submitter: submitter}]
}

// loadTaints fills taints from the datadir; a missing file is fine. Entries
// written before per-submitter taints existed (epoch:aid) load as
// whole-application taints.
func (n *Node) loadTaints() {
	if n.taintFile == "" {
		return
	}
	raw, err := os.ReadFile(n.taintFile)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	if err != nil {
		log.Warnw("tainted applications: cannot read, starting empty", "path", n.taintFile, "err", err)
		return
	}
	var keys []string
	if err := json.Unmarshal(raw, &keys); err != nil {
		log.Warnw("tainted applications: malformed file, starting empty", "path", n.taintFile, "err", err)
		return
	}
	for _, k := range keys {
		if tk, ok := parseTaintKey(k); ok {
			n.taints[tk] = true
		}
	}
	if len(n.taints) > 0 {
		log.Infow("tainted applications loaded", "count", len(n.taints))
	}
}

// saveTaints writes taints atomically; failures only cost a repeated search.
func (n *Node) saveTaints() {
	if n.taintFile == "" {
		return
	}
	keys := make([]string, 0, len(n.taints))
	for tk := range n.taints {
		keys = append(keys, tk.String())
	}
	raw, err := json.Marshal(keys)
	if err != nil {
		return
	}
	tmp := n.taintFile + ".tmp"
	err = os.MkdirAll(filepath.Dir(n.taintFile), 0o700)
	if err == nil {
		err = os.WriteFile(tmp, raw, 0o600)
	}
	if err == nil {
		err = os.Rename(tmp, n.taintFile)
	}
	if err != nil {
		log.Warnw("tainted applications: cannot persist", "path", n.taintFile, "err", err)
	}
}

// String is the persisted form: epoch:aid:submitter. The two-part
// epoch:aid spelling only survives on disk from whole-application taints
// written before the per-submitter rule.
func (tk taintKey) String() string {
	s := hex.EncodeToString(tk.epoch[:]) + ":" + hex.EncodeToString(tk.aid[:])
	if tk.submitter != (common.Address{}) {
		s += ":" + hex.EncodeToString(tk.submitter[:])
	}
	return s
}

func parseTaintKey(s string) (taintKey, bool) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 && len(parts) != 3 {
		return taintKey{}, false
	}
	e, err1 := hex.DecodeString(parts[0])
	a, err2 := hex.DecodeString(parts[1])
	if err1 != nil || err2 != nil || len(e) != 12 || len(a) != 32 {
		return taintKey{}, false
	}
	var tk taintKey
	copy(tk.epoch[:], e)
	copy(tk.aid[:], a)
	if len(parts) == 3 {
		sub, err := hex.DecodeString(parts[2])
		if err != nil || len(sub) != common.AddressLength {
			return taintKey{}, false
		}
		tk.submitter = common.BytesToAddress(sub)
	}
	return tk, true
}
