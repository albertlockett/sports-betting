vendor:
	go mod vendor

build:
	go build -o bin/main main.go

docker:
	docker build -t sports-betting:latest .

docker-service:
	docker build -f Dockerfile-service -t sports-betting-service:latest .