@echo off

SET nodejs_version=%1
SET tonosse_version=%2
SET arango_version=%3
SET qserver=%4
SET nginx_version=1.18.0
SET tonossePath=%userprofile%\tonse

:: Caddy
mkdir %tonossePath%
mkdir %tonossePath%\caddy
cd %tonossePath%\caddy
curl -O https://github.com/caddyserver/caddy/releases/download/v2.4.0-beta.2/caddy_2.4.0-beta.2_windows_amd64.zip
tar xf caddy_2.4.0-beta.2_windows_amd64.zip
curl -O https://raw.githubusercontent.com/INTONNATION/tonos-se-installers/master/tonsectl/caddy/Caddyfile

:: Release downloading
cd %tonossePath%
curl -LJO https://github.com/INTONNATION/tonos-se-installers/releases/download/%tonosse_version%/tonos-se-windows.tar
tar xf tonos-se-windows.tar
move tonos-se/docker/ton-live/web/ $tonossePath/caddy/web/
DEL /Q tonos-se-windows.tar

:: ArangoDB

curl -O https://download.arangodb.com/arangodb37/Community/Windows/ArangoDB3-%arango_version%_win64.zip
tar xf ArangoDB3-%arango_version%_win64.zip
DEL /Q ArangoDB3-%arango_version%_win64.zip
move ArangoDB3-%arango_version%_win64 %tonossePath%\arangodb
mkdir %tonossePath%\arangodb\etc
mkdir %tonossePath%\arangodb\var\lib\arangodb3
mkdir %tonossePath%\arangodb\initdb.d
curl https://raw.githubusercontent.com/INTONNATION/tonos-se-installers/master/tonsectl/arangodb/config -o %tonossePath%\arangodb\etc\config
curl https://raw.githubusercontent.com/tonlabs/tonos-se/%tonosse_version%/docker/arango/initdb.d/upgrade-arango-db.js -o %tonossePath%\arangodb\initdb.d\upgrade-arango-db.js

:: TON node

mkdir %tonossePath%\node
cd %tonossePath%\node

curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/ton-node/blockchain.conf.json
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/$tonosse_version/docker/ton-node/ton-node.conf.json
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/%tonosse_version%/docker/ton-node/log_cfg.yml
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/%tonosse_version%/docker/ton-node/private-key
curl -O https://raw.githubusercontent.com/tonlabs/tonos-se/%tonosse_version%/docker/ton-node/pub-key

move %tonossePath%\ton_node_startup.exe %tonossePath%\node\ton_node_startup.exe

:: Graph QL

mkdir %tonossePath%\graphql
cd %tonossePath%\graphql
curl -O https://nodejs.org/dist/v%nodejs_version%/node-v%nodejs_version%-win-x64.zip
tar xf node-v%nodejs_version%-win-x64.zip
del /Q node-v%nodejs_version%-win-x64.zip
move node-v%nodejs_version%-win-x64 nodejs

move %tonossePath%\%qserver% %tonossePath%\graphql\

set PATH=%PATH%;%tonossePath%\graphql\nodejs

tar xf %qserver%
curl -o %tonossePath%\graphql\package\.env https://raw.githubusercontent.com/INTONNATION/tonos-se-installers/master/tonsectl/graphql/.env

npm install %qserver% --production
del /Q /S %qserver%
