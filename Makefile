.DEFAULT_GOAL := default
.PHONY: all helm
BINARY_NAME=o1

get:
	go get ./...

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/o1/server/main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-client-linux cmd/o1/client/main.go

build_darwin:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin cmd/o1/server/main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-client-darwin cmd/o1/client/main.go

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
	helm dependency update ./helm
	helm install o1 ./helm
	kubectl rollout status statefulset/o1-etcd
	kubectl rollout status statefulset/o1

helm_delete:
	helm delete o1 || true
	kubectl delete pvc  data-o1-etcd-0 || true

minikube:
	minikube start --memory 2048 --cpus 2
	eval $(minikube docker-env)

minikube_stop:
	minikube stop

etcd:
	docker stop etcd || true
	docker rm etcd || true
	docker run -d  -p 2379:2379 --name etcd -v /usr/share/ca-certificates/:/etc/ssl/certs quay.io/coreos/etcd:v3.5.1 /usr/local/bin/etcd -advertise-client-urls http://0.0.0.0:2379 -listen-client-urls http://0.0.0.0:2379

etcd_logs:
	docker logs etcd

port_forward:
	kubectl port-forward service/o1 6969:6969

default: docker helm_delete helm port_forward