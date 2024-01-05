# HA

高可用系统

### feature

- 支持多种类型数据库高可用
- 健康探测/主备切换/拓扑管理/recovery/
- 多种recovery策略
- raft，保证HA自身的分布式高可用

### usage

#### 准备工作

使用mysql创建表：

1. 调用脚本创建表，数据表位于pkg/resource/sql
2. 调用cmd/gen/generate.go生成db/model

#### deploy

```shell
#start 3 nodes
$ ./ha --raft_conf=conf/orchestrator-7.conf.json 
$ ./ha --raft_conf=conf/orchestrator-8.conf.json
$ ./ha --raft_conf=conf/orchestrator-9.conf.json
$ go install github.com/Jille/raftadmin/cmd/raftadmin@latest
$ raftadmin localhost:10007 add_voter localhost:10008 localhost:10008 0
```
