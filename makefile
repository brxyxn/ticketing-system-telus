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

# database
connect-database:
	docker run --rm -it -d -p 5432:5432 --env-file .env --name dbpostgres --network technical-test_database --volume technical-test-temp:/var/lib/postgresql/data postgres:14.2

psql:
	docker exec -it dbpostgres psql -U telus_user telus_db