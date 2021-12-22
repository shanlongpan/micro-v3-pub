GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install github.com/golang/protobuf/proto
	go install github.com/golang/protobuf/protoc-gen-go
	# 下面的安装如果报错，单独安装一下，执行 go install github.com/asim/go-micro/cmd/protoc-gen-micro/v3@latest
	go install github.com/asim/go-micro/cmd/protoc-gen-micro/v3
#	go install github.com/asim/go-micro/v3/cmd/protoc-gen-openapi

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/newmicro.proto

.PHONY: proto
proto:
	#protoc --proto_path=./idl --micro_out=./idl/grpc --go_out=./idl/grpc helloworld.proto
	protoc --proto_path=. --micro_out=. --go_out=:. idl/helloworld.proto
	
.PHONY: build
