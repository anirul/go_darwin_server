package darwin_service

import (
	"context"
	"errors"
	"sync"
	"time"

	"google.golang.org/grpc/peer"

	math "github.com/anirul/go_darwin_server/darwin_math"
	proto "github.com/anirul/go_darwin_server/darwin_proto"
)

type DarwinService struct {
	proto.UnimplementedDarwinServiceServer
	mu        sync.Mutex
	world     *proto.WorldDatabase
	peerChars map[string]string
}

func NewDarwinService(worldDatabase *proto.WorldDatabase) *DarwinService {
	return &DarwinService{world: worldDatabase}
}

func (s *DarwinService) Planet() *proto.Element {
	for _, p := range s.world.Elements {
		if p.Physic.Radius > 100 {
			return p
		}
	}
	return nil
}

func (s *DarwinService) Ping(
	ctx context.Context, pingRequest *proto.PingRequest) (
	*proto.PingResponse, error) {
	return &proto.PingResponse{
		Value:           pingRequest.GetValue(),
		Time:            math.TimeSecondNow(),
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
	p, ok := peer.FromContext(ctx)
	if !ok {
		return &proto.CreateCharacterResponse{ReturnEnum: proto.ReturnEnum_RETURN_ERROR},
			errors.New("invalid peer")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, character := range s.world.Characters {
		if character.Name == createRequest.Name {
			return &proto.CreateCharacterResponse{ReturnEnum: proto.ReturnEnum_RETURN_REJECTED},
				errors.New("name already taken")
		}
	}
	if !math.IsInColor(createRequest.Color, s.world.PlayerParameter.ColorParameters) {
		return &proto.CreateCharacterResponse{ReturnEnum: proto.ReturnEnum_RETURN_REJECTED},
			errors.New("not a valid color")
	}
	planet := s.Planet()
	randomNormalizeVec3 := math.RandomNormalizeVec3()
	newCharacter := &proto.Character{}
	newCharacter.Name = createRequest.Name
	newCharacter.Color = createRequest.Color
	newCharacter.Physic.Radius = math.RadiusFromVolume(s.world.PlayerParameter.StartMass)
	newCharacter.Physic.Mass = s.world.PlayerParameter.StartMass
	newCharacter.Physic.Position = math.Times(
		randomNormalizeVec3,
		planet.Physic.Radius+s.world.PlayerParameter.DropHeight)
	newCharacter.StatusEnum = proto.StatusEnum_STATUS_LOADING
	s.peerChars[p.Addr.String()] = createRequest.Name
	s.world.Characters = append(s.world.Characters, newCharacter)
	return &proto.CreateCharacterResponse{ReturnEnum: proto.ReturnEnum_RETURN_OK}, nil
}

func (s *DarwinService) Update(
	req *proto.UpdateRequest, stream proto.DarwinService_UpdateServer) error {
	for {
		time.Sleep(time.Millisecond * 100)
		s.mu.Lock()
		err := stream.Send(&proto.UpdateResponse{
			Characters: s.world.Characters,
			Elements:   s.world.Elements,
			Time:       math.TimeSecondNow(),
		})
		s.mu.Unlock()
		if err != nil {
			return err
		}
	}
}
