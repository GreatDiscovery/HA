package raft

import (
	"context"
	"fmt"
	"ha/pkg/log"
	"ha/pkg/runtime"
)

func SetUp(ctx context.Context) {
	log.G(ctx).Info("setting up raft")
	hostname := runtime.HostName
	fmt.Println("hostname=", hostname)
}
