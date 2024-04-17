package darwin_service

import (
	"context"
	"time"

	darwin "github.com/anirul/go_darwin_server/darwin_math"
	proto "github.com/anirul/go_darwin_server/darwin_proto"
)

type DarwinService struct {
	proto.UnimplementedDarwinServiceServer
	world *proto.WorldDatabase
}

func NewDarwinService(worldDatabase *proto.WorldDatabase) *DarwinService {
	return &DarwinService{world: worldDatabase}
}

func (s *DarwinService) Ping(
	ctx context.Context, pingRequest *proto.PingRequest) (
	*proto.PingResponse, error) {
	return &proto.PingResponse{
		Value:           pingRequest.GetValue(),
		Time:            float64(time.Now().UnixNano()) / 1e9,
		PlayerParameter: s.world.PlayerParameter,
	}, nil
}

func (s *DarwinService) ReportInGame(
	ctx context.Context, reportRequest *proto.ReportInGameRequest) (
	*proto.ReportInGameResponse, error) {
	return nil, nil
}

func (s *DarwinService) CreateCharacter(
	ctx context.Context, createRequest *proto.CreateCharacterRequest) (
	*proto.CreateCharacterResponse, error) {
	newCharacter := proto.Character{}
	newCharacter.Name = createRequest.Name
	newCharacter.Color = createRequest.Color
	newCharacter.Physic = &proto.Physic{}
	newCharacter.Physic.Radius = 1.0
	newCharacter.Physic.Mass = 1.0
	return nil, nil
}

func (s *DarwinService) Update(
	req *proto.UpdateRequest, stream proto.DarwinService_UpdateServer) error {
	for {
		time.Sleep(time.Millisecond * 100)
		if err := stream.Send(&proto.UpdateResponse{
			Characters: s.world.Characters,
			Elements:   s.world.Elements,
			Time:       darwin.TimeSecondNow(),
		}); err != nil {
			return err
		}
	}
}
