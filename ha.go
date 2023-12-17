package main

import (
	"context"
	"flag"
	"ha/pkg/config"
	"ha/pkg/cron"
	"ha/pkg/discovery"
	"ha/pkg/http"
	"ha/pkg/log"
	"ha/pkg/register"
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
	todo := context.TODO()
	configuration, err := config.NewConfiguration(*raftConf)
	if err != nil {
		os.Exit(1)
	}

	logManager := log.NewLogManager(configuration)
	logManager.SetUp()
	log.G(root).Infof("configuration: %v", configuration)

	cronManager := cron.NewCronManager(configuration)
	cronManager.SetUp()
	defer cronManager.Shutdown()

	// 2. register self in period
	processManager := register.NewProcessManager(configuration)
	processManager.SetUp()
	_ = processManager.Registering(todo)

	// 3. starting discovery
	discoveryManager := discovery.NewDiscoveryManager(configuration)
	discoveryManager.Discovery(root)

	// 4. web server start
	http.Setup(configuration.ListenAddress)
}
