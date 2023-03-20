build:
	go build -o bin/server app/main.go
run: build
	./bin/server
tidy:
	go mod tidy
docker-build:
	docker build -t albingeorge/commently-service .
docker-run: docker-build
	docker run -p 8080:8080 -it albingeorge/commently-service /app/server
