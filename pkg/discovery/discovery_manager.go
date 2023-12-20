package discovery

import (
	"context"
	"ha/pkg/config"
	"ha/pkg/log"
	"ha/pkg/raft"
	"ha/pkg/service"
)

type discoveryManager struct {
	RaftEnabled bool
	RaftBind    string
}

func NewDiscoveryManager(configuration config.Configuration) service.DiscoveryManager {
	return &discoveryManager{
		RaftEnabled: configuration.RaftEnabled,
		RaftBind:    configuration.RaftBind,
	}
}

func (d *discoveryManager) Discovery(ctx context.Context) error {
	log.G(ctx).Info("starting auto discovery")
	if !d.RaftEnabled {
		log.G(ctx).Info("raft disabled")
		return nil
	}
	raftConfig := raft.Config{
		RaftBind: d.RaftBind,
	}
	err := raft.SetUp(ctx, raftConfig)
	if err != nil {
		return err
	}
	return nil
}
