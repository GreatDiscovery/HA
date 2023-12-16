package register

import (
	"HA/pkg/config"
	"HA/pkg/log"
	"HA/pkg/service"
	"context"
)

type processManager struct {
}

func NewProcessManager(configuration config.Configuration) service.ProcessManager {
	return &processManager{}
}

func (p *processManager) SetUp() {
	log.G(context.TODO()).Infof("starting processManager")
}

func (p *processManager) Registering(ctx context.Context) error {
	return nil
}
