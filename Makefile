# Docker
docker-run:
	docker run --rm -it --name telus-container telus-tickets:latest

docker-build:
	docker build -t telus-tickets:latest .

## Compose
compose-up:
	docker-compose up --build

## Swarm for secrets
secrets-init:
	docker swarm init
