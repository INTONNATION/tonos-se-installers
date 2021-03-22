#!/bin/bash

nodejs_version="v12.21.0"
tonosse_version="0.25.0"
arango_version="3.7.9"

mkdir -p ~/tonse
tonossePath="$HOME/tonse"

# Download tonosse and extract TON node and Graph binaries

cd $tonossePath
curl -LJO https://github.com/INTONNATION/tonos-se-installers/releases/download/tonos-se-v-$tonosse_version/tonos-se-v-$tonosse_version.tgz
mv tonos-se-v-$tonosse_version.tgz tonos-se-v-$tonosse_version.tar
tar xf tonos-se-v-$tonosse_version.tar
rm tonos-se-v-$tonosse_version.tar

# Arango DB
curl -O https://download.arangodb.com/arangodb37/Community/MacOSX/arangodb3-macos-$arango_version.tar.gz

tar xf arangodb3-macos-$arango_version.tar.gz
rm arangodb3-macos-$arango_version.tar.gz

mv arangodb3-macos-$arango_version $tonossePath/arangodb
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

mv $tonossePath/ton_node_startup $tonossePath/node/
chmod +x $tonossePath/node/ton_node_startup

# Graph QL

mkdir -p $tonossePath/graphql
cd $tonossePath/graphql

curl "https://nodejs.org/dist//latest-v12.x/node-${nodejs_version:-$(wget -qO- https://nodejs.org/dist/latest/ | sed -nE 's|.*>node-(.*)\.pkg</a>.*|\1|p')}.pkg" > "$HOME/Downloads/node-latest.pkg" && sudo installer -store -pkg "$HOME/Downloads/node-latest.pkg" -target "/"
qserver=`ls $tonossePath | grep ton-q-server`
mv $tonossePath/$qserver $tonossePath/graphql/
npm config set registry="http://registry.npmjs.org"
npm install $qserver --production
tar xf $tonossePath/graphql/$qserver
rm -rf $tonossePath/graphql/$qserver