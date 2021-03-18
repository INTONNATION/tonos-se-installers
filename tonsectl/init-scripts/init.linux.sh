#!/bin/bash

nodejs_version="16.16.0"
tonosse_version="0.25.0"
arango_version="3.7.9"

mkdir -p ~/tonse
tonossePath="$HOME/tonse"

# Download tonosse and extract TON node and Graph binaries

cd $tonossePath
curl -O https://github.com/INTONNATION/tonos-se-installers/releases/download/tonos-se-v-$tonosse_version/tonos-se-v-$tonosse_version.tgz
tar xf tonos-se-v-$tonosse_version.tgz
rm tonos-se-v-$tonosse_version.tgz

# Arango DB
curl -O https://download.arangodb.com/arangodb37/Community/Linux/arangodb3-linux-$arango_version.tar.gz

tar xf arangodb3-linux-$arango_version.tar.gz
rm arangodb3-linux-$arango_version.tar.gz

mv arangodb3-linux-$arango_version $tonossePath/arangodb
mkdir -p $tonossePath/arangodb/etc
mkdir -p $tonossePath/arangodb/var/lib/arangodb3
mkdir -p $tonossePath/arangodb/initdb.d/

cp arangodb/config $tonossePath/arangodb/etc/config
curl https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/arango/initdb.d/upgrade-arango-db.js -o $tonossePath/arangodb/initdb.d/upgrade-arango-db.js

# TON node

mkdir -p $tonossePath/node
cd $tonossePath/node

curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/cfg
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/log_cfg.yml
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/private-key
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/pub-key

mv $tonossePath/ton_node_startup $tonossePath/node/
chmod +x $tonossePath/node/ton_node_startup

# Graph QL

mkdir -p $tonossePath/graphql
cd $tonossePath/graphql

curl -O https://nodejs.org/dist/v$nodejs_version/node-v$nodejs_version-linux-x64.tar.xz 
tar -xf node-v$nodejs_version-linux-x64.tar.xz
rm node-v$nodejs_version-linux-x64.tar.xz
mv ./node-v$nodejs_version-linux-x64/ ./nodejs
qserver=`ls $tonossePath | grep ton-q-server`
mv $tonossePath/$qserver $tonossePath/graphql/
./nodejs/bin/npm install $qserver
