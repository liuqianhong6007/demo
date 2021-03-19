package room

import (
	"context"
	"encoding/json"
	"testing"

	"google.golang.org/grpc"
)

func toJson(o interface{}) string {
	buf, _ := json.Marshal(o)
	return string(buf)
}

func newGrpcClient() RoomClient {
	addr := "127.0.0.1:8811"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := NewRoomClient(conn)
	return client
}

func BenchmarkImplRoomServer_GetRoomList(b *testing.B) {
	client := newGrpcClient()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			in := &GetRoomListIn{}
			_, err := client.GetRoomList(ctx, in)
			if err != nil {
				panic(err)
			}
		}
	})
}
