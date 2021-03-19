package room

import (
	"context"
)

type ImplRoomServer struct {
}

func (ImplRoomServer) GetRoomList(context.Context, *GetRoomListIn) (*GetRoomListOut, error) {
	room := &RoomInfo{
		RoomId:   1,
		RoomName: "default_room",
		MaxNum:   10,
		CurNum:   1,
		ModeId:   0,
		MapId:    "default_map",
		State:    RoomInfo_WAITING_GAME,
	}

	return &GetRoomListOut{
		Rooms: []*RoomInfo{room},
	}, nil
}
func (ImplRoomServer) mustEmbedUnimplementedRoomServer() {}
