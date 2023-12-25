package discovery

import (
	"encoding/json"
	"github.com/hashicorp/raft"
	_const "ha/pkg/const"
	"ha/pkg/log"
	"io"
)

// finite state machine
type fsm Store

func (f *fsm) Apply(l *raft.Log) interface{} {
	var c storeCommand
	err := json.Unmarshal(l.Data, &c)
	if err != nil {
		log.G(_const.TODO).Errorf("failed to unmarshal command: %s", err.Error())
	}
	return store.applier.ApplyCommand(c.Op, c.Value)
}

// Snapshot persist data
func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	//TODO implement me
	panic("implement me")
}

func (f *fsm) Restore(snapshot io.ReadCloser) error {
	//TODO implement me
	panic("implement me")
}
