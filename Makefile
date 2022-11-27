.PHONY: up db-up db-down dev build
up:
	go install github.com/cosmtrek/air@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/rakyll/gotest@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install entgo.io/ent/cmd/ent@latest

db-up: # dbを起動する
	docker-compose up -d

db-down: # dbを停止してデータを削除する
	docker-compose down --rmi all
	rm -r db/data

dev:
	air

build:
	go build .

# lint:
# 	golangci-lint run ./...