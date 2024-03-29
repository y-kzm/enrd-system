##### Docker Build #####
IMG_CONTROLLER = yykzm/ubuntu:20.04-controller
IMG_AGENT = yykzm/ubuntu:20.04-agent

docker-build-controller:
	docker build -t $(IMG_CONTROLLER) -f docker/Dockerfile.controller .
docker-build-agent:
	docker build -t $(IMG_AGENT) -f docker/Dockerfile.agent .
rmi:
	docker rmi $(IMG_CONTROLLER)
	docker rmi $(IMG_AGENT)

##### Go Build #####
controller:
	go build -o ./bin/controller cmd/controller/main.go

agent:
	make cgo
	go build -o ./bin/agent cmd/agent/main.go

clean:
	cd ./pkg/tool/igi-ptr && make clean
	rm -rf ./bin

cgo:
	cd ./pkg/tool/igi-ptr && make clean && make all
 
proto:
	protoc --go_out=./api --go_opt=paths=source_relative --go-grpc_out=./api --go-grpc_opt=paths=source_relative --proto_path=./api/protos ./api/protos/*.proto	