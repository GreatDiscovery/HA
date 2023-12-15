package main

import (
	"HA/pkg/config"
	"HA/pkg/discovery"
	"HA/pkg/http"
	"HA/pkg/log"
	"HA/pkg/register"
	"context"
	"flag"
	"os"
)

var (
	raftConf = flag.String("raft_conf", "conf/raft.conf.json", "raft config filepath")
	//discovery = flag.Bool("discovery", true, "auto discovery mode")
)

func main() {
	root := context.TODO()
	// 1. init config
	flag.Parse()
	raftConfig, err := config.NewConfiguration(*raftConf)
	if err != nil {
		os.Exit(1)
	}

	logManager := log.NewLogManager(raftConfig)
	logManager.SetUp()
	log.G(root).Infof("raftConfig: %v", raftConfig)

	// 2. register self in period
	processManager := register.NewProcessManager(raftConfig)
	processManager.SetUp()

	// 3. starting discovery
	discoveryManager := discovery.NewDiscoveryManager(raftConfig)
	discoveryManager.Discovery(root)

	// 4. web server start
	http.Setup(raftConfig.ListenAddress)
}
