package discovery

import "github.com/hashicorp/raft"

type Store struct {
	raftDir  string
	raftBind string
	raft     *raft.Raft
}

func NewStore(raftDir string, raftBind string) *Store {
	return &Store{
		raftDir:  raftDir,
		raftBind: raftBind,
	}
}

// SetUpRaft raft core ini method
func (store *Store) SetUpRaft(peerNodes []string) error {
	return nil
}
