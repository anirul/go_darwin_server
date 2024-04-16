package server

import (
	darwin "github.com/anirul/go_darwin_server/darwin"
)

type DarwinService struct {
	darwin.UnimplementedDarwinServiceServer
}

func NewDarwinService() *DarwinService {
	return &DarwinService{}
}
