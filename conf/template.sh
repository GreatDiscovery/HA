#!/bin/bash

generate_raft_conf() {
  local NODE_ID="$1"

  echo "setting up orchestrator-$NODE_ID..."

  mkdir -p "/tmp/orchestrator-$NODE_ID"

  cp raft.conf.json "orchestrator-$NODE_ID.conf.json"
  sed -i '' -e "s/NODE_ID_PLACEHOLDER/$NODE_ID/g" orchestrator-"$NODE_ID".conf.json
}

#generate 3 raft configs
for id in 7 8 9 ; do
  generate_raft_conf "$id"
done