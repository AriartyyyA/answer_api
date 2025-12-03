.PHONY: build run docker-build docker-up docker-down clean

build:
	go build -o main ./cmd/app

run:
	go run ./cmd/app

docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

clean:
	rm -f main
