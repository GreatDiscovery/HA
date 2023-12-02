package HA

import (
	"HA/pkg/config"
	"flag"
	"fmt"
)

var (
	raftConf = flag.String("raft_conf", "conf/raft.conf.json", "raft config filepath")
)

func main() {
	// 1. init config
	flag.Parse()
	raftConfig := config.NewConfiguration(*raftConf)
	fmt.Println(raftConfig)

	// 2. run raft
}
