install:
	go mod tidy
	go mod vendor

up: build
	docker-compose build --no-cache
	docker-compose up -d

down:
	docker-compose down -v

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/tikad ./cmd/tikad