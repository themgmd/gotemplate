.PHONY: dev run build

build: 
	go build -o .bin/main cmd/main/app.go

up:
	docker-compose up -d --build

dev:
	go run cmd/main/app.go

run: build
	.bin/main