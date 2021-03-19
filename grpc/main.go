package main

import (
	"flag"
	"net"

	"google.golang.org/grpc"

	"github.com/liuqianhong6007/demo/grpc/room"
)

func main() {
	addr := flag.String("addr", "0.0.0.0:8811", "grpc address")
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	room.RegisterRoomServer(grpcServer, room.ImplRoomServer{})

	grpcServer.Serve(lis)
}
