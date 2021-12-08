.DEFAULT_GOAL := default
.PHONY: all helm
BINARY_NAME=o1

build:
	go get ./...
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/o1/server/server.go

build_darwin:
	go get ./...
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin cmd/o1/server/server.go

run:
	./bin/${BINARY_NAME}-linux

clean:
	go clean
	rm bin/${BINARY_NAME}-linux

docker:
	docker build . -t loqutus/o-1
	docker push loqutus/o-1

helm:
	helm install o1 ./helm
	helm delete o1

default: build