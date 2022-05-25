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

