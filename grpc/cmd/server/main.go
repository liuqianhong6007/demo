package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/liuqianhong6007/demo/grpc/cmd/server/config"
	"github.com/liuqianhong6007/demo/grpc/cmd/server/module/room"
)

func main() {
	lis, err := net.Listen("tcp", config.GrpcAddr())
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	room.RegisterRoomServer(grpcServer, room.ImplRoomServer{})

	grpcServer.Serve(lis)
}
