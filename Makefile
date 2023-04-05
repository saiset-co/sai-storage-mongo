up:
	docker-compose -f ./microservices/docker-compose.yml up -d

down:
	docker-compose -f ./microservices/docker-compose.yml down --remove-orphans

build:
	make service
	make docker

service:
		cd ./src/saiStorageMongo && go mod tidy && go build -o ../../microservices/saiStorageMongo/build/sai-storage
		cp ./src/saiStorageMongo/config.json ./microservices/saiStorageMongo/build/config.json
docker:
	docker-compose -f ./microservices/docker-compose.yml up -d --build

log:
	docker-compose -f ./microservices/docker-compose.yml logs -f

logs:
	docker-compose -f ./microservices/docker-compose.yml logs -f sai-storage

