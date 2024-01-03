package discovery

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
)

type SnapshotCreatorApplier interface {
	GetData() (data []byte, err error)
	Restore(rc io.ReadCloser) error
}

type SnapshotData struct {
}

type SnapshotDataCreatorApplier struct{}

func NewSnapshotData() *SnapshotData {
	return &SnapshotData{}
}

// GetData read from db
func (s *SnapshotDataCreatorApplier) GetData() (data []byte, err error) {
	snapshotData := CreateSnapshotData()
	b, err := json.Marshal(snapshotData)
	if err != nil {
		return b, err
	}
	var buf bytes.Buffer
	// compress data
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(b); err != nil {
		return b, err
	}
	if err := zw.Close(); err != nil {
		return b, err
	}
	return buf.Bytes(), nil
}

// Restore write to db from file
func (s *SnapshotDataCreatorApplier) Restore(rc io.ReadCloser) error {
	snapshotData := NewSnapshotData()
	zr, err := gzip.NewReader(rc)
	if err != nil {
		return err
	}
	err = json.NewDecoder(zr).Decode(snapshotData)
	if err != nil {
		return err
	}
	return nil
}

func NewSnapshotDataCreatorApplier() *SnapshotDataCreatorApplier {
	generator := &SnapshotDataCreatorApplier{}
	return generator
}

func CreateSnapshotData() *SnapshotData {
	snapshotData := NewSnapshotData()
	// todo gain data
	return snapshotData
}
