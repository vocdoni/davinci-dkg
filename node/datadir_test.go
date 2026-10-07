package node

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

func TestDeploymentDir(t *testing.T) {
	c := qt.New(t)
	manager := common.HexToAddress("0xC6Fb38c42ed3FB35D363a702218746d5C7Da36BF")
	c.Assert(deploymentDir("/data", 100, manager), qt.Equals,
		filepath.Join("/data", "100-0xc6fb38c42ed3fb35d363a702218746d5c7da36bf"))
	c.Assert(deploymentDir("", 100, manager), qt.Equals, "")
}

func epochWithPrefix(prefix uint32, nonce uint64) [12]byte {
	var id [12]byte
	binary.BigEndian.PutUint32(id[:4], prefix)
	binary.BigEndian.PutUint64(id[4:], nonce)
	return id
}

// Legacy state moves into the deployment directory only when its epoch ids
// carry this deployment's prefix; another deployment's state stays put.
func TestMigrateLegacyState(t *testing.T) {
	c := qt.New(t)
	const own, other = uint32(0xaabbccdd), uint32(0x11223344)
	datadir := t.TempDir()
	dir := deploymentDir(datadir, 100, common.HexToAddress("0x01"))
	dealer := common.HexToAddress("0xd1")

	legacy := &contributionCache{dir: contributionCacheDir(datadir)}
	ownEpochID, otherEpochID := epochWithPrefix(own, 1), epochWithPrefix(other, 1)
	legacy.Put(ownEpochID, dealer, []byte("own"))
	legacy.Put(otherEpochID, dealer, []byte("other"))

	ownTaint := taintKey{epoch: ownEpochID, aid: [32]byte{1}, submitter: dealer}.String()
	otherTaint := taintKey{epoch: otherEpochID, aid: [32]byte{2}}.String()
	kept := taintKey{epoch: epochWithPrefix(own, 2), aid: [32]byte{3}}.String()
	c.Assert(writeTaintKeys(taintPath(datadir), []string{ownTaint, otherTaint}), qt.IsNil)
	c.Assert(writeTaintKeys(taintPath(dir), []string{kept}), qt.IsNil)

	migrateLegacyState(datadir, dir, own)

	moved := &contributionCache{dir: contributionCacheDir(dir)}
	got, ok := moved.Get(ownEpochID, dealer)
	c.Assert(ok, qt.IsTrue)
	c.Assert(string(got), qt.Equals, "own")
	_, ok = moved.Get(otherEpochID, dealer)
	c.Assert(ok, qt.IsFalse)
	_, ok = legacy.Get(ownEpochID, dealer)
	c.Assert(ok, qt.IsFalse)
	got, ok = legacy.Get(otherEpochID, dealer)
	c.Assert(ok, qt.IsTrue)
	c.Assert(string(got), qt.Equals, "other")

	keys, err := readTaintKeys(taintPath(dir))
	c.Assert(err, qt.IsNil)
	c.Assert(keys, qt.DeepEquals, []string{kept, ownTaint})
	keys, err = readTaintKeys(taintPath(datadir))
	c.Assert(err, qt.IsNil)
	c.Assert(keys, qt.DeepEquals, []string{otherTaint})

	// A second run is a no-op.
	migrateLegacyState(datadir, dir, own)
	keys, err = readTaintKeys(taintPath(dir))
	c.Assert(err, qt.IsNil)
	c.Assert(keys, qt.DeepEquals, []string{kept, ownTaint})

	// Once every legacy entry has moved, the legacy files go.
	migrateLegacyState(datadir, dir, other)
	_, err = os.Stat(contributionCacheDir(datadir))
	c.Assert(os.IsNotExist(err), qt.IsTrue)
	_, err = os.Stat(taintPath(datadir))
	c.Assert(os.IsNotExist(err), qt.IsTrue)
	_, ok = moved.Get(otherEpochID, dealer)
	c.Assert(ok, qt.IsTrue)
}
