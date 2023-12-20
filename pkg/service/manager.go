package service

import (
	"context"
	"errors"
	"sync"
)

// Manager manage your own process, like log/monitor/health etc.
// The methods should be thread-safe.
type Manager interface {
	// SetUp manager start only once
	SetUp()
	// Shutdown manager
	Shutdown()
}

type MangerConfig struct {
	sync.Mutex
	Initialized bool
}

var UnInitialized = errors.New("service not initialized")

type ProcessManager interface {
	Manager
	// Registering will continuously update the node_health table showing that the current process is still running.
	Registering(ctx context.Context) error
}

type DiscoveryManager interface {
	// Discovery starts an asynchronous infinite discovery process where instances are
	// periodically investigated and their status captured, and long since unseen instances are
	// purged and forgotten.
	Discovery(ctx context.Context) error
}

type LogManager interface {
	Manager
}

type CronManager interface {
	Manager
}
