#!/bin/bash

mkdir ~/tonse
tonossePath="$HOME/tonse"

wget https://download.arangodb.com/arangodb37/Community/Linux/arangodb3-linux-3.7.9.tar.gz

tar xf arangodb3-linux-3.7.9.tar.gz
rm -f arangodb3-linux-3.7.9.tar.gz

mv arangodb3-linux-3.7.9 $tonossePath/arangodb
mkdir -p $tonossePath/arangodb/etc
mkdir -p $tonossePath/arangodb/var/lib/arangodb3
mkdir -p $tonossePath/arangodb/initdb.d/

cp arangodb/config $tonossePath/arangodb/etc/config
wget https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/arango/initdb.d/upgrade-arango-db.js -O $tonossePath/arangodb/initdb.d/upgrade-arango-db.js
