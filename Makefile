.PHONY: up up-db down-db conn-db dev build
up:
	go install github.com/cosmtrek/air@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/rakyll/gotest@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install entgo.io/ent/cmd/ent@latest

up-db: # dbを起動する
	docker-compose up -d

down-db: # dbを停止してデータを削除する
	docker-compose down --rmi all
	rm -rf database/data

conn-db: # dbコンテナに接続
	PGPASSWORD=postgres psql -h localhost -p 4432 -U postgres -d postgres

dev:
	air

build:
	go build .

# lint:
# 	golangci-lint run ./...