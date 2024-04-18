package darwin_service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc/peer"

	math "github.com/anirul/go_darwin_server/darwin_math"
	proto "github.com/anirul/go_darwin_server/darwin_proto"
)

type CharacterHit struct {
	character proto.Character
	hit       string
}

type DarwinService struct {
	proto.UnimplementedDarwinServiceServer
	mu            sync.Mutex
	world         *proto.WorldDatabase
	peerChars     map[string]string
	potentialHits []CharacterHit
}

func NewDarwinService(worldDatabase *proto.WorldDatabase) *DarwinService {
	return &DarwinService{world: worldDatabase, peerChars: make(map[string]string)}
}

func (s *DarwinService) Planet() *proto.Element {
	for _, p := range s.world.Elements {
		if p.Name == "earth" && p.TypeEnum == proto.TypeEnum_TYPE_GROUND {
			return p
		}
	}
	return nil
}

func (s *DarwinService) Peer(ctx context.Context) (*peer.Peer, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("invalid peer")
	}
	return p, nil
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
	p, err := s.Peer(ctx)
	if err != nil {
		return &proto.ReportInGameResponse{}, nil
	}
	_, exists := s.peerChars[p.Addr.String()]
	if !exists {
		return &proto.ReportInGameResponse{}, nil
	}
	found := false
	for i, character := range s.world.Characters {
		if character.Name == reportRequest.Name {
			s.world.Characters[i].Physic = reportRequest.Physic
			s.world.Characters[i].StatusEnum = reportRequest.StatusEnum
			s.world.Characters[i].SpecialEffectBoost = reportRequest.SpecialEffectBoost
			if reportRequest.PotentialHit != "" {
				s.potentialHits = append(
					s.potentialHits,
					CharacterHit{*character, reportRequest.PotentialHit},
				)
			}
			found = true
			return &proto.ReportInGameResponse{}, nil
		}
	}
	if !found {
		return &proto.ReportInGameResponse{}, nil
	}
	return &proto.ReportInGameResponse{}, nil
}

func (s *DarwinService) CreateCharacter(
	ctx context.Context, createRequest *proto.CreateCharacterRequest) (
	*proto.CreateCharacterResponse, error) {
	p, err := s.Peer(ctx)
	if err != nil {
		return &proto.CreateCharacterResponse{ReturnEnum: proto.ReturnEnum_RETURN_ERROR}, nil
	}
	for _, character := range s.world.Characters {
		if character.Name == createRequest.Name {
			return &proto.CreateCharacterResponse{ReturnEnum: proto.ReturnEnum_RETURN_REJECTED}, nil
		}
	}
	if !math.IsInColor(createRequest.Color, s.world.PlayerParameter.ColorParameters) {
		return &proto.CreateCharacterResponse{ReturnEnum: proto.ReturnEnum_RETURN_REJECTED}, nil
	}
	planet := s.Planet()
	if planet == nil {
		return &proto.CreateCharacterResponse{ReturnEnum: proto.ReturnEnum_RETURN_ERROR}, nil
	}
	randomNormalizeVec3 := math.RandomNormalizeVec3()
	newCharacter := &proto.Character{}
	newCharacter.Name = createRequest.Name
	newCharacter.Color = createRequest.Color
	newCharacter.Physic = &proto.Physic{}
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
		fmt.Print(".")
	}
}
