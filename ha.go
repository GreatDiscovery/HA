package main

import (
	"HA/pkg/config"
	"HA/pkg/log"
	"HA/pkg/register"
	"context"
	"flag"
	"os"
)

var (
	raftConf  = flag.String("raft_conf", "conf/raft.conf.json", "raft config filepath")
	discovery = flag.Bool("discovery", true, "auto discovery mode")
)

func main() {
	root := context.TODO()
	// 1. init config
	flag.Parse()
	_ = log.SetFormat(log.TextFormat)
	raftConfig, err := config.NewConfiguration(*raftConf)
	if err != nil {
		os.Exit(1)
	}
	log.G(root).Infof("%#v", raftConfig)
	if raftConfig.Debug {
		_ = log.SetLevel(log.Debug)
	}

	// 2. register self in period
	_, err = register.NewProcessManager(raftConfig)
	if err != nil {
		log.G(root).Error("registerManager init failed")
		os.Exit(1)
	}

	// 3. web server start

}
