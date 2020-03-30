dependencies:
	dep ensure

build_searchd_from_container:
	dep ensure
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./bin/searchd ./cmd/searchd

run_build_container:
	docker build --file=Dockerfile.searchd.build -t starship-build .
	docker run -it -v $(shell pwd):/root/go/src/github.com/dannav/starship/services starship-build /bin/sh -c 'cd /root/go/src/github.com/dannav/starship/services; make build_searchd_from_container;'

build_tikad:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/tikad ./cmd/tikad

build: run_build_container build_tikad

test:
	# NOTICE
	# ======
	# ensure that annoy is installed before running tests
	# run: `sudo pip install annoy`
	go test ./cmd/searchd/...