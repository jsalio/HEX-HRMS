.PHONY: build clean run run-webui test

build:
	go build -o main cmd/main.go

test:
	go test ./...

clean:
	rm main

run:
	go run cmd/main.go

run-webui:
	cd infra/hrms-web && npm run dev

run-all:
	go run cmd/main.go & cd infra/hrms-web && npm run start

