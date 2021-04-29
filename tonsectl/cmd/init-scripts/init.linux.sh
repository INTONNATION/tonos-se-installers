#!/bin/bash

nodejs_version="$1"
tonosse_version="$2"
arango_version="$3"

set -e
tonossePath="$HOME/tonse"

# Caddy

mkdir -p $tonossePath/caddy
cd $tonossePath/caddy

curl -s -LJ -o caddy.tar.gz https://github.com/caddyserver/caddy/releases/download/v2.4.0-beta.2/caddy_2.4.0-beta.2_linux_amd64.tar.gz
tar -zxf caddy.tar.gz 
chmod +x caddy
sudo setcap cap_net_bind_service=ep $tonossePath/caddy/caddy
rm caddy.tar.gz
curl -O https://raw.githubusercontent.com/INTONNATION/tonos-se-installers/master/tonsectl/caddy/Caddyfile

# Download tonosse and extract TON node and Graph binaries

cd $tonossePath
curl -LJO https://github.com/INTONNATION/tonos-se-installers/releases/download/$tonosse_version/tonos-se-linux.tgz
tar xf tonos-se-linux.tgz
cp -r tonos-se/docker/ton-live/web/ $tonossePath/caddy/web/
rm tonos-se-linux.tgz

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

curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/ton-node/blockchain.conf.json
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/ton-node/ton-node.conf.json
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/ton-node/log_cfg.yml
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/ton-node/private-key
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/ton-node/pub-key

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
PATH=$PATH:$tonossePath/graphql/nodejs/bin/
tar xf $tonossePath/graphql/$qserver
rm -rf $tonossePath/graphql/$qserver 
cd $tonossePath/graphql/package
npm install --production
