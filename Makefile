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

#protoc:
#	protoc --go_out=./api --go_opt=paths=source_relative --go-grpc_out=./api --go-grpc_opt=paths=source_relative --proto_path=./api/protos ./api/protos/*.proto	
