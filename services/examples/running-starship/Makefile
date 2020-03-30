download_model:
	@mkdir -p ./use
	python ./create-serving-use.py

setup:
	(cd ../../ && docker build -f Dockerfile.searchd -t searchd .) # build and tag docker images
	(cd ../../ && docker build -f Dockerfile.tikad -t tikad .) # copy all docker files to this directory
	docker-compose build --no-cache # build services defined in docker-compose.yaml

up:
	docker-compose up -d && docker-compose logs -tf # run Starship services

down:
	docker-compose down -v # stop Starship services and delete application data
