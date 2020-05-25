.PHONY: build run dc-up

DOCKERDIR = docker

build:
	go build main.go

run:
	go run cmd/app/main.go

dc-up:
	cd "${DOCKERDIR}" && docker-compose up

dc-down:
	cd "${DOCKERDIR}" && docker-compose down
