# backend
build:
	go build -o ./backend/bin/server ./backend/cmd/main.go

run: build
	./backend/bin/server

watch:
	reflex -s -r '\.go$$' make run
