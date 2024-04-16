package server

import (
	proto "github.com/anirul/go_darwin_server/darwin_proto"
)

type DarwinService struct {
	proto.UnimplementedDarwinServiceServer
}

func NewDarwinService() *DarwinService {
	return &DarwinService{}
}
