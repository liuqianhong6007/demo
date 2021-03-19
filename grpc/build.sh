#!/bin/bash

go_out=`pwd`/room
proto_dir=`pwd`/proto

protoc --proto_path=${proto_dir} --go_out=${go_out} --go-grpc_out=${go_out}  ${proto_dir}/room.proto

# 编译 go 代码
export GOBIN=`pwd`/bin
go install
