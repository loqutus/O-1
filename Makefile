.DEFAULT_GOAL := default
.PHONY: all helm
BINARY_NAME=o1

get:
	go get ./...

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/o1/server/server.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-client-linux cmd/o1/client/client.go

build_darwin:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin cmd/o1/server/server.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-client-darwin cmd/o1/client/client.go

run:
	./bin/${BINARY_NAME}-linux

run_darwin:
	./bin/${BINARY_NAME}-darwin

test:
	go test ./...

clean:
	go clean
	rm bin/${BINARY_NAME}-linux

docker:
	docker build . -t loqutus/o-1
	docker push loqutus/o-1
	docker build . -f Dockerfile-client -t loqutus/o-1-client
	docker push loqutus/o-1-client

docker_run:
	docker stop o1 || true
	docker rm o1 || true
	docker run -d -p 6969:6969 --name o1 loqutus/o-1

docker_logs:
	docker logs o1

helm:
	helm install o1 ./helm

helm_delete:
	helm delete o1

minikube:
	minikube start --memory 2048 --cpus 2
	eval $(minikube docker-env)

minikube_stop:
	minikube stop

etcd:
	docker stop etcd || true
	docker rm etcd || true
	docker run -d  -p 2379:2379 --name etcd -v /usr/share/ca-certificates/:/etc/ssl/certs quay.io/coreos/etcd:v3.5.1 -listen-client-urls http://0.0.0.0:2379

default: get minikube docker docker_run etcd test minikube_stop