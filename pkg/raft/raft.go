package raft

import (
	"context"
	"fmt"
	"ha/pkg/log"
	"ha/pkg/runtime"
)

type Config struct {
	RaftBind string // host:port
}

func SetUp(ctx context.Context, config Config) {
	log.G(ctx).Info("setting up raft")
	hostname := runtime.HostName
	fmt.Println("hostname=", hostname)
}
