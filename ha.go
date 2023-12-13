package main

import (
	"HA/pkg/config"
	"HA/pkg/log"
	"flag"
	"fmt"
	"os"
)

var (
	raftConf = flag.String("raft_conf", "conf/raft.conf.json", "raft config filepath")
)

func main() {
	// 1. init config
	flag.Parse()
	_ = log.SetFormat(log.TextFormat)
	raftConfig, err := config.NewConfiguration(*raftConf)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(raftConfig)

	// 2. run raft
}
