package service

import "context"

// define runtime service interface

// RuntimeService interface should be implemented by a container runtime.
// The methods should be thread-safe.
type RuntimeService interface {
	ProcessManager
}

type ProcessManager interface {
	// Registering will continuously update the node_health table showing that the current process is still running.
	Registering(ctx context.Context) error
}
