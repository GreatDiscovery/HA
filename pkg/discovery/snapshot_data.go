package discovery

import "io"

type SnapshotCreatorApplier interface {
	GetData() (data []byte, err error)
	Restore(rc io.ReadCloser) error
}

type SnapshotDataCreatorApplier struct{}

func (s *SnapshotDataCreatorApplier) GetData() (data []byte, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *SnapshotDataCreatorApplier) Restore(rc io.ReadCloser) error {
	//TODO implement me
	panic("implement me")
}

func NewSnapshotDataCreatorApplier() *SnapshotDataCreatorApplier {
	generator := &SnapshotDataCreatorApplier{}
	return generator
}
