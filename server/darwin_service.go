package server

import (
	"context"
	"time"

	proto "github.com/anirul/go_darwin_server/darwin_proto"
)

type DarwinService struct {
	proto.UnimplementedDarwinServiceServer
	playerParameter *proto.PlayerParameter
}

func NewDarwinService(playerParam *proto.PlayerParameter) *DarwinService {
	return &DarwinService{playerParameter: playerParam}
}

func (s *DarwinService) Ping(ctx context.Context, pingRequest *proto.PingRequest) (*proto.PingResponse, error) {
	return &proto.PingResponse{
		Value:           pingRequest.GetValue(),
		Time:            float64(time.Now().UnixNano()) / 1e9,
		PlayerParameter: s.playerParameter,
	}, nil
}
