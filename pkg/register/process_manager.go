package register

import (
	"context"
	"ha/pkg/config"
	"ha/pkg/log"
	"ha/pkg/service"
)

type processManager struct {
}

func NewProcessManager(configuration config.Configuration) service.ProcessManager {
	return &processManager{}
}

func (p *processManager) SetUp() {
	log.G(context.TODO()).Infof("starting processManager")
}

func (p *processManager) Shutdown() {
	//TODO implement me
	panic("implement me")
}

func (p *processManager) Registering(ctx context.Context) error {
	return nil
}
