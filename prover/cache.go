package prover

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CachedArtifact describes one serialized artifact written into the local cache.
type CachedArtifact struct {
	Hash string
	Path string
	Size int64
}

// WriteArtifact stores a byte slice under its SHA-256 hash in dir.
func WriteArtifact(dir string, content []byte) (*CachedArtifact, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("create artifact dir %s: %w", dir, err)
	}
	sum := sha256.Sum256(content)
	hash := hex.EncodeToString(sum[:])
	path := filepath.Join(dir, hash)
	if err := os.WriteFile(path, content, 0o644); err != nil {
		return nil, fmt.Errorf("write artifact %s: %w", path, err)
	}
	return &CachedArtifact{
		Hash: hash,
		Path: path,
		Size: int64(len(content)),
	}, nil
}

// WriteWriterArtifact serializes any value exposing WriteTo into the artifact cache.
func WriteWriterArtifact(dir string, value interface {
	WriteTo(io.Writer) (int64, error)
},
) (*CachedArtifact, error) {
	var buf bytes.Buffer
	if _, err := value.WriteTo(&buf); err != nil {
		return nil, fmt.Errorf("serialize artifact: %w", err)
	}
	return WriteArtifact(dir, buf.Bytes())
}
