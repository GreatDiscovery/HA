package discovery

import (
	"context"
	"ha/pkg/config"
	"ha/pkg/log"
	"ha/pkg/service"
)

type discoveryManager struct {
}

func NewDiscoveryManager(configuration config.Configuration) service.DiscoveryManager {
	return discoveryManager{}
}

func (d discoveryManager) Discovery(ctx context.Context) {
	log.G(ctx).Info("starting auto discovery")
}
