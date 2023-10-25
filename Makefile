all: tidy format build test

tidy:
	go mod tidy -v

format:
	go fmt

build:
	go build -v ./...

test:
	go test -v ./...

update:
	go get -u all
