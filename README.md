# HA
高可用系统

### feature

- 支持多种类型数据库高可用
- 健康探测/主备切换/拓扑管理/recovery/
- 多种recovery策略
- raft，保证HA自身的分布式高可用

### usage
#### DB 
使用mysql
1. 调用脚本创建表
2. 调用cmd/gen/generate.go生成db/model