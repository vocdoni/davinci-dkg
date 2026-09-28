package node

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/log"
)

// Everything the node persists (contribution calldata, taints) belongs to one
// deployment, so it lives under <datadir>/<chainid>-<manager>. A node that
// moves to another deployment, because a release changed the preset or the
// operator changed --manager, starts from an empty directory, and moving back
// finds the old state again.

// deploymentDir returns the state directory of one deployment ("" for no
// datadir, which disables persistence).
func deploymentDir(datadir string, chainID uint64, manager common.Address) string {
	if datadir == "" {
		return ""
	}
	return filepath.Join(datadir, fmt.Sprintf("%d-%s", chainID, strings.ToLower(manager.Hex())))
}

// migrateLegacyState moves state written straight into datadir, before
// per-deployment directories existed, into dir. Epoch ids start with the
// manager's EPOCH_PREFIX, so only this deployment's entries move; anything
// else stays where it is and is never read.
func migrateLegacyState(datadir, dir string, prefix uint32) {
	if datadir == "" || dir == "" {
		return
	}
	migrateLegacyContributions(contributionCacheDir(datadir), contributionCacheDir(dir), prefix)
	migrateLegacyTaints(taintPath(datadir), taintPath(dir), prefix)
}

// ownEpoch reports whether epoch id belongs to the deployment with prefix.
func ownEpoch(epoch [12]byte, prefix uint32) bool {
	return binary.BigEndian.Uint32(epoch[:4]) == prefix
}

func migrateLegacyContributions(from, to string, prefix uint32) {
	entries, err := os.ReadDir(from)
	if err != nil {
		return
	}
	moved := 0
	for _, e := range entries {
		raw, err := hex.DecodeString(e.Name())
		if !e.IsDir() || err != nil || len(raw) != 12 || !ownEpoch([12]byte(raw), prefix) {
			continue
		}
		dst := filepath.Join(to, e.Name())
		if _, err := os.Stat(dst); err == nil {
			continue
		}
		err = os.MkdirAll(to, 0o700)
		if err == nil {
			err = os.Rename(filepath.Join(from, e.Name()), dst)
		}
		if err != nil {
			log.Warnw("contribution cache: cannot move legacy entry", "from", from, "epoch", e.Name(), "err", err)
			continue
		}
		moved++
	}
	if moved > 0 {
		log.Infow("contribution cache: moved legacy entries into the deployment directory", "epochs", moved, "dir", to)
	}
	_ = os.Remove(from) // only succeeds once empty
}

func migrateLegacyTaints(from, to string, prefix uint32) {
	legacy, err := readTaintKeys(from)
	if err != nil || len(legacy) == 0 {
		return
	}
	var own, rest []string
	for _, k := range legacy {
		if tk, ok := parseTaintKey(k); ok && ownEpoch(tk.epoch, prefix) {
			own = append(own, k)
		} else {
			rest = append(rest, k)
		}
	}
	if len(own) == 0 {
		return
	}
	current, err := readTaintKeys(to)
	if err != nil {
		log.Warnw("tainted applications: cannot read, legacy entries not moved", "path", to, "err", err)
		return
	}
	merged := current
	for _, k := range own {
		if !slices.Contains(merged, k) {
			merged = append(merged, k)
		}
	}
	if err := writeTaintKeys(to, merged); err != nil {
		log.Warnw("tainted applications: cannot persist, legacy entries not moved", "path", to, "err", err)
		return
	}
	if len(rest) == 0 {
		err = os.Remove(from)
	} else {
		err = writeTaintKeys(from, rest)
	}
	if err != nil {
		// Harmless: the entries are already in the new file and merge again next time.
		log.Warnw("tainted applications: cannot rewrite legacy file", "path", from, "err", err)
	}
	log.Infow("tainted applications: moved legacy entries into the deployment directory", "count", len(own), "path", to)
}
