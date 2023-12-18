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
}

func NewDiscoveryManager(configuration config.Configuration) service.DiscoveryManager {
	return &discoveryManager{
		RaftEnabled: configuration.RaftEnabled,
	}
}

func (d *discoveryManager) Discovery(ctx context.Context) {
	log.G(ctx).Info("starting auto discovery")
	if !d.RaftEnabled {
		log.G(ctx).Info("raft disabled")
		return
	}
	raft.SetUp(ctx)
}
