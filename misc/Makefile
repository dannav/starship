build_tikad:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../services/bin/tikad ../services/cmd/tikad

up: build_tikad
	docker build -t searchd ../services -f ../services/Dockerfile.searchd
	docker build -t tikad ../services -f ../services/Dockerfile.tikad
	docker-compose build --no-cache
	docker-compose up -d

down:
	docker-compose down -v

deploy_bins:
	./deploy/push.sh