package node

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/vocdoni/davinci-dkg/log"
)

// Tainted applications are remembered on disk so a restarted node does not
// repeat the 2^50 search that condemned them: a fleet restart would otherwise
// run one full-core search per node per poisoned application at once.

func taintPath(datadir string) string {
	if datadir == "" {
		return ""
	}
	return filepath.Join(datadir, "tainted-apps.json")
}

// loadTaints fills taintedApps from the datadir; a missing file is fine.
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
		if ak, ok := parseAppKey(k); ok {
			n.taintedApps[ak] = true
		}
	}
	if len(n.taintedApps) > 0 {
		log.Infow("tainted applications loaded", "count", len(n.taintedApps))
	}
}

// saveTaints writes taintedApps atomically; failures only cost a repeated search.
func (n *Node) saveTaints() {
	if n.taintFile == "" {
		return
	}
	keys := make([]string, 0, len(n.taintedApps))
	for ak := range n.taintedApps {
		keys = append(keys, hex.EncodeToString(ak.epoch[:])+":"+hex.EncodeToString(ak.aid[:]))
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

func parseAppKey(s string) (appKey, bool) {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return appKey{}, false
	}
	e, err1 := hex.DecodeString(parts[0])
	a, err2 := hex.DecodeString(parts[1])
	if err1 != nil || err2 != nil || len(e) != 12 || len(a) != 32 {
		return appKey{}, false
	}
	var ak appKey
	copy(ak.epoch[:], e)
	copy(ak.aid[:], a)
	return ak, true
}
