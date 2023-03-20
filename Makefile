build:
	go build -o bin/main app/main.go
run: build
	./bin/main
tidy:
	go mod tidy
