package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"

	"github.com/liuqianhong6007/demo/grpc/cmd/client/config"
	"github.com/liuqianhong6007/demo/grpc/cmd/client/module/room"
)

func toJson(o interface{}) string {
	buf, _ := json.Marshal(o)
	return string(buf)
}

func main() {
	conn, err := grpc.Dial(config.GrpcAddr(), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := room.NewRoomClient(conn)

	ctx := context.Background()

	// 模拟请求
	in := &room.GetRoomListIn{}
	out, err := client.GetRoomList(ctx, in)
	if err != nil {
		panic(err)
	}
	fmt.Println(toJson(out))
}
