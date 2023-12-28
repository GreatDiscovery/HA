package discovery

import "io"

type SnapshotCreatorApplier interface {
	GetData() (data []byte, err error)
	Restore(rc io.ReadCloser) error
}

type SnapshotDataCreatorApplier struct {
}

func NewSnapshotDataCreatorApplier() *SnapshotDataCreatorApplier {
	generator := &SnapshotDataCreatorApplier{}
	return generator
}
