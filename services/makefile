install:
	go mod tidy
	go mod vendor

up: build
	docker-compose build --no-cache
	docker-compose up -d

down:
	docker-compose down -v

build_searchd_from_container:
	dep ensure
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./bin/searchd ./cmd/searchd

run_build_container:
	docker build --file=dockerfile.searchd.build -t starship-build .
	docker run -it -v $(shell pwd):/root/go/src/github.com/dannav/starship/services starship-build /bin/sh -c 'cd /root/go/src/github.com/dannav/starship/services; make build_searchd_from_container;'

build: run_build_container
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/tikad ./cmd/tikad

deploy_bins:
	./deploy/push.sh