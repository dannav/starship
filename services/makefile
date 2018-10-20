install:
	go mod tidy
	go mod vendor

build:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/searchd ./searchd

up: build
	docker-compose up --build -d

stop:
	docker-compose stop