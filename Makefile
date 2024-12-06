build: 
	docker-compose build
start:
	docker-compose up --build
push:
	docker build -t ghcr.io/database-sabermetrics/cyber-esteban:latest .
	docker push ghcr.io/database-sabermetrics/cyber-esteban:latest

