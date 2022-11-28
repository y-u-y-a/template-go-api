include .env.development

.PHONY: up down dev build
init:
	go install github.com/cosmtrek/air@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/rakyll/gotest@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install entgo.io/ent/cmd/ent@latest

up: # dbを起動する
	docker-compose up -d

down: # dbを停止してデータを削除する
	docker-compose down
	rm -rf database/data

dev: # サーバー起動
	air

build:
	go build .

.PHONY: gen-migration migrate-up migrate-down conn-db

gen-migration: # スキーマ変更による差分からマイグレーションファイルを生成
	cd database; go run ./diff.go

migrate-up: # golang-migrate
	migrate -database "$(POSTGRESQL_URL)" -path database/migrations up
migrate-down: # golang-migrate
	migrate -database "$(POSTGRESQL_URL)" -path database/migrations down -all

conn-db: # ローカル環境からdbコンテナに接続
	PGPASSWORD=postgres psql -h localhost -p 4432 -U postgres -d postgres