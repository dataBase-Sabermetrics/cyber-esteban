build: 
	docker build -t cyber-esteban .

run: 
	docker run -p 8080:8080 cyber-esteban
