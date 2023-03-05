# export NODE_OPTIONS=--openssl-legacy-provider

npm cache clean --force
docker builder prune
docker compose up -d
<!-- docker-compose up go -->

# install後に処理
go clean -cache
go mod tidy
go mod init kakeibo2

docker-compose exec kakeibo2-api bash
docker-compose exec mysql bash

source ~/.bashrc

## dockerビルド
docker-compose up
docker-compose build --no-cache
docker-compose up

docker-compose exec go bash

go run main.go

## DB接続
docker-compose exec mysql bash
docker-compose exec mysql bash
mysql -u root -p
