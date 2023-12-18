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
	"os/signal"
	"syscall"
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

	// 2. setup managers
	logManager := log.NewLogManager(configuration)
	logManager.SetUp()
	log.G(root).Infof("configuration: %+v", configuration)

	// 3. process hook
	go registerSignals(todo)

	cronManager := cron.NewCronManager(configuration)
	cronManager.SetUp()
	defer cronManager.Shutdown()

	processManager := register.NewProcessManager(configuration)
	processManager.SetUp()
	_ = processManager.Registering(todo)

	// 4. starting discovery
	discoveryManager := discovery.NewDiscoveryManager(configuration)
	discoveryManager.Discovery(root)

	// 4. web server start
	http.Setup(configuration.ListenAddress)
}

func registerSignals(ctx context.Context) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGHUP)
	signal.Notify(c, syscall.SIGTERM)

	for sig := range c {
		switch sig {
		case syscall.SIGHUP:
			log.G(ctx).Info("Received SIGHUP. Reloading configuration")
			//todo config.reload()
		case syscall.SIGTERM:
			log.G(ctx).Info("Received SIGTERM. Shutting down")
			os.Exit(0)
		}
	}
}
