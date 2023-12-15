package service

import (
	"context"
	"sync"
)

type SetUp interface {
	// SetUp manager start only once
	SetUp()
}

type SetUpConfig struct {
	sync.Mutex
	Initialized bool
}

// Manager manage your own process, like log/monitor/health etc.
// The methods should be thread-safe.
type Manager interface {
	ProcessManager
	DiscoveryManager
	LogManager
}

type ProcessManager interface {
	// Registering will continuously update the node_health table showing that the current process is still running.
	Registering(ctx context.Context) error
}

type DiscoveryManager interface {
	// Discovery starts an asynchronous infinite discovery process where instances are
	// periodically investigated and their status captured, and long since unseen instances are
	// purged and forgotten.
	Discovery(ctx context.Context)
}

type LogManager interface {
	SetUp
}
