build:
	go build -o bin/jorat cmd/main.go
run:
	./bin/jorat
serve:
	clear;
	make build
	make run
