# backend
build:
	go build -o ./backend/bin/server ./backend/cmd/main.go

run: build
	./backend/bin/server

watch:
	reflex -s -r '\.go$$' make run



# frontend
react-build:
	cd frontend && npm run build && cd ..

react-start:
	cd frontend && npm start && cd ..

# full app
start-app:
	make react-start && make watch




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
