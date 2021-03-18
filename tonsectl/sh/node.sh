#!/bin/bash

mkdir ~/tonse
tonossePath="$HOME/tonse"

mkdir -p $tonossePath/node
cd tonossePath/node

curl -O https://github.com/INTONNATION/tonos-se-installers/releases/download/tonos-se-v-0.25.0/tonos-se-v-0.25.0.tgz
tar xf tonos-se-v-0.25.0.tgz
rm -f tonos-se-v-0.25.0.tgz

curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/cfg
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/log_cfg.yml
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/private-key
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/pub-key

chmod +x $tonossePath/node/ton_node_startup
