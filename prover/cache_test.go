package prover

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestWriteArtifact(t *testing.T) {
	c := qt.New(t)

	dir := t.TempDir()
	content := []byte("artifact-bytes")

	result, err := WriteArtifact(dir, content)
	c.Assert(err, qt.IsNil)
	c.Assert(result.Path, qt.Equals, filepath.Join(dir, result.Hash))

	expected := sha256.Sum256(content)
	c.Assert(result.Hash, qt.Equals, hex.EncodeToString(expected[:]))

	stored, err := os.ReadFile(result.Path)
	c.Assert(err, qt.IsNil)
	c.Assert(stored, qt.DeepEquals, content)
}

type testWriterTo struct {
	data []byte
}

func (w testWriterTo) WriteTo(buf io.Writer) (int64, error) {
	n, err := buf.Write(w.data)
	return int64(n), err
}

func TestWriteWriterArtifact(t *testing.T) {
	c := qt.New(t)

	dir := t.TempDir()
	result, err := WriteWriterArtifact(dir, testWriterTo{data: []byte("writer-artifact")})
	c.Assert(err, qt.IsNil)
	c.Assert(result.Hash != "", qt.IsTrue)

	stored, err := os.ReadFile(result.Path)
	c.Assert(err, qt.IsNil)
	c.Assert(string(stored), qt.Equals, "writer-artifact")
}
