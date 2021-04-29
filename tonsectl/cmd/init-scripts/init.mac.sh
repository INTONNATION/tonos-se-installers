#!/bin/bash

nodejs_version="$1"
tonosse_version="$2"
arango_version="$3"

tonossePath="$HOME/tonse"

# Caddy

mkdir -p $tonossePath/caddy
cd $tonossePath/caddy

curl -LJ -o caddy.tar.gz https://github.com/caddyserver/caddy/releases/download/v2.4.0-beta.2/caddy_2.4.0-beta.2_mac_amd64.tar.gz
tar -xvf caddy.tar.gz 
chmod +x caddy
rm caddy.tar.gz
curl -O https://raw.githubusercontent.com/INTONNATION/tonos-se-installers/master/tonsectl/caddy/Caddyfile

# Download tonosse and extract TON node and Graph binaries

cd $tonossePath
curl -LJO https://github.com/INTONNATION/tonos-se-installers/releases/download/$tonosse_version/tonos-se-darwin.tar
tar xf tonos-se-darwin.tar
cp -r tonos-se/docker/ton-live/web/ $tonossePath/caddy/web/
rm tonos-se-darwin.tar

# Arango DB
curl -O https://download.arangodb.com/arangodb37/Community/MacOSX/arangodb3-macos-$arango_version.tar.gz

tar xf arangodb3-macos-$arango_version.tar.gz
rm arangodb3-macos-$arango_version.tar.gz

mv arangodb3-macos-$arango_version $tonossePath/arangodb
mkdir -p $tonossePath/arangodb/etc
mkdir -p $tonossePath/arangodb/var/lib/arangodb3
mkdir -p $tonossePath/arangodb/initdb.d/

curl https://raw.githubusercontent.com/INTONNATION/tonos-se-installers/master/tonsectl/arangodb/config -o $tonossePath/arangodb/etc/config
curl https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/arango/initdb.d/upgrade-arango-db.js -o $tonossePath/arangodb/initdb.d/upgrade-arango-db.js

# TON node

mkdir -p $tonossePath/node
cd $tonossePath/node

curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/ton-node/blockchain.conf.json
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/ton-node/ton-node.conf.json
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/ton-node/log_cfg.yml
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/ton-node/private-key
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/ton-node/pub-key

mv $tonossePath/ton_node_startup $tonossePath/node/
chmod +x $tonossePath/node/ton_node_startup

# Graph QL

mkdir -p $tonossePath/graphql
cd $tonossePath/graphql

curl -O https://nodejs.org/dist/v$nodejs_version/node-v$nodejs_version.pkg && sudo installer -store -pkg "node-v$nodejs_version.pkg" -target "/"
qserver=`ls $tonossePath | grep ton-q-server`
mv $tonossePath/$qserver $tonossePath/graphql/
npm config set registry="http://registry.npmjs.org"
npm install $qserver --production
tar xf $tonossePath/graphql/$qserver
rm -rf $tonossePath/graphql/$qserver
