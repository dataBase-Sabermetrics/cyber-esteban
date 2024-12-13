SHA := $(shell openssl rand -hex 8)

.PHONY: push

build: 
	docker-compose -f .docker/compose.yml build
start:
	docker-compose -f .docker/compose.yml up --build
push:
	docker build -t ghcr.io/database-sabermetrics/cyber-esteban:$(SHA) -t ghcr.io/database-sabermetrics/cyber-esteban:latest -f .docker/Dockerfile .
	docker push ghcr.io/database-sabermetrics/cyber-esteban:$(SHA)
