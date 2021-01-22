#!/bin/bash

function Exec() {
  echo "# $*"
  $*
  if [ $? != 0 ];then
    exit
  fi
}

function generateCode() {
  go_out=$1
  proto_dir=$2

  if [ x"${go_out}" = "x" ];then
    echo "not define go_out"
    return 1
  fi

  if [ x"${proto_dir}" = "x" ];then
    echo "not define proto_dir"
    return 1
  fi

  if [ -d ${go_out} ];then
    rm -rf ${go_out}/*.pb.go
  else
    mkdir ${go_out}
  fi

  protoc --proto_path=${proto_dir} --go_out=${go_out} --go-grpc_out=${go_out}  ${proto_dir}/room.proto
}

function gobuild() {
    module_path=$1
    exec_name=$2

    if [ ! -d ${module_path} ];then
      echo "module_path not exist: ${module_path}"
      return 1
    fi

    if [ -z "${exec_name}" ];then
      echo "go exec name not define"
      return 1
    fi

    cd ${module_path} || return 1
    go install $module_path
}

GO_OUT=`pwd`/cmd
PROTO_DIR=`pwd`/proto

# 生成 grpc 代码
Exec generateCode ${GO_OUT}/server $PROTO_DIR
Exec generateCode ${GO_OUT}/client $PROTO_DIR

# 编译 go 代码
export GOBIN=`pwd`/bin
if [ ! -d $GOBIN ];then
  mkdir $GOBIN
fi

Exec gobuild ${GO_OUT}/server server
Exec gobuild ${GO_OUT}/client client


