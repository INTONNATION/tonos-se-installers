#!/bin/bash

nodejs_version="14.16.0"
tonosse_version="0.25.0"
arango_version="3.7.9"

set -e

mkdir -p ~/tonse
tonossePath="$HOME/tonse"

# Download tonosse and extract TON node and Graph binaries

cd $tonossePath
curl -LJO https://github.com/INTONNATION/tonos-se-installers/releases/download/tonos-se-v-$tonosse_version/tonos-se-v-$tonosse_version.tgz
tar xf tonos-se-v-$tonosse_version.tgz
rm tonos-se-v-$tonosse_version.tgz

# Arango DB
curl -O https://download.arangodb.com/arangodb37/Community/Linux/arangodb3-linux-$arango_version.tar.gz

tar xf arangodb3-linux-$arango_version.tar.gz
rm arangodb3-linux-$arango_version.tar.gz

rm $tonossePath/arangodb -rf
mv arangodb3-linux-$arango_version $tonossePath/arangodb
mkdir -p $tonossePath/arangodb/etc
mkdir -p $tonossePath/arangodb/var/lib/arangodb3
mkdir -p $tonossePath/arangodb/initdb.d/

curl https://raw.githubusercontent.com/INTONNATION/tonos-se-installers/master/tonsectl/arangodb/config -o $tonossePath/arangodb/etc/config
curl https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/arango/initdb.d/upgrade-arango-db.js -o $tonossePath/arangodb/initdb.d/upgrade-arango-db.js

# TON node

mkdir -p $tonossePath/node
cd $tonossePath/node

curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/cfg
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/log_cfg.yml
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/private-key
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/pub-key

rm $tonossePath/node/ton_node_startup -f
mv $tonossePath/ton_node_startup $tonossePath/node/
chmod +x $tonossePath/node/ton_node_startup

# Graph QL

mkdir -p $tonossePath/graphql
cd $tonossePath/graphql

curl -O https://nodejs.org/dist/v$nodejs_version/node-v$nodejs_version-linux-x64.tar.xz 
tar -xf node-v$nodejs_version-linux-x64.tar.xz
rm node-v$nodejs_version-linux-x64.tar.xz
rm -rf ./nodejs
mv ./node-v$nodejs_version-linux-x64/ ./nodejs
qserver=`ls $tonossePath | grep ton-q-server`
echo $qserver
rm -rf $tonossePath/graphql/$qserver
mv $tonossePath/$qserver $tonossePath/graphql/
PATH=$PATH:$tonossePath/nodejs/bin/
sudo mkdir /usr/lib/node_modules -p
sudo chown -R $USER /usr/lib/node_modules
./nodejs/bin/npm install $qserver --production -g
rm -rf $tonossePath/graphql/$qserver 
