package register

import (
	"HA/pkg/config"
	"HA/pkg/service"
	"context"
)

type processManager struct {
}

func NewProcessManager(configuration config.Configuration) (service.ProcessManager, error) {
	return processManager{}, nil
}

func (p processManager) Registering(ctx context.Context) error {
	return nil
}
