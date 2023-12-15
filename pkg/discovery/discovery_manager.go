package discovery

import (
	"HA/pkg/config"
	"HA/pkg/log"
	"HA/pkg/service"
	"context"
)

type discoveryManager struct {
}

func NewDiscoveryManager(configuration config.Configuration) service.DiscoveryManager {
	return discoveryManager{}
}

func (d discoveryManager) Discovery(ctx context.Context) {
	log.G(ctx).Info("starting auto discovery")
}
