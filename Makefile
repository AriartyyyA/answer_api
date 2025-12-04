.PHONY: build run docker-build docker-up docker-down clean

run:
	go run ./cmd/app

docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

