package discovery

import (
	"github.com/hashicorp/raft"
	"io"
)

// finite state machine
type fsm Store

func (*fsm) Apply(log *raft.Log) interface{} {
	//TODO implement me
	panic("implement me")
}

func (*fsm) Snapshot() (raft.FSMSnapshot, error) {
	//TODO implement me
	panic("implement me")
}

func (*fsm) Restore(snapshot io.ReadCloser) error {
	//TODO implement me
	panic("implement me")
}
