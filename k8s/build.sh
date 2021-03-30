#!/bin/bash

function gen_grpc() {
  mkdir -p watcher/protocol
  protoc -I watcher/protos watcher.proto --go_out=watcher/protocol --go-grpc_out=watcher/protocol
}

function build_image() {
    docker build -t lqha.xyz/k8s-test:latest .
}

function print_usage() {
    echo "usage: ./build.sh [gen|build]"
    exit
}

cmd=$1

case $cmd in
  "gen")
  gen_grpc
  ;;

  "build")
  build_image
  ;;

  *)
  print_usage
  ;;
esac



