build: 
	docker build -t cyber-esteban:latest .
run: 
	docker rm -f cyber-esteban || true
	docker run -p 8080:8080 --name cyber-esteban cyber-esteban:latest

