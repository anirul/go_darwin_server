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
	ptr       *proto.Character
	physic    proto.Physic
	hitTarget string
}

type HitType int

const (
	HitTypeNone HitType = iota
	HitTypeElement
	HitTypeCharacter
)

type PeerClient struct {
	characterName string
	lastSeen      float64
}

type DarwinService struct {
	proto.UnimplementedDarwinServiceServer
	mu            sync.Mutex
	world         *proto.WorldDatabase
	peerChars     map[string]PeerClient
	potentialHits []CharacterHit
}

func NewDarwinService(worldDatabase *proto.WorldDatabase) *DarwinService {
	return &DarwinService{world: worldDatabase, peerChars: make(map[string]PeerClient)}
}

func (s *DarwinService) planet() *proto.Element {
	for _, p := range s.world.Elements {
		if p.Name == "earth" && p.TypeEnum == proto.TypeEnum_TYPE_GROUND {
			return p
		}
	}
	return nil
}

func (s *DarwinService) peer(ctx context.Context) (*peer.Peer, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("invalid peer")
	}
	return p, nil
}

func (s *DarwinService) removeDisconnected() {
	for key, value := range s.peerChars {
		if math.TimeSecondNow()-value.lastSeen > 10 {
			for i, character := range s.world.Characters {
				if character.Name == value.characterName {
					s.world.Characters = append(
						s.world.Characters[:i], s.world.Characters[i+1:]...)
					delete(s.peerChars, key)
					break
				}
			}
		}
	}
}

func (s *DarwinService) physicAndColor(name string) (HitType, *proto.Physic, *proto.Vector3) {
	for _, character := range s.world.Characters {
		if character.Name == name {
			return HitTypeCharacter, character.Physic, character.Color
		}
	}
	for _, element := range s.world.Elements {
		if element.Name == name {
			return HitTypeElement, element.Physic, element.Color
		}
	}
	return HitTypeNone, nil, nil
}

func (s *DarwinService) checkHit() {
	for _, ch := range s.potentialHits {
		ht, physic, color := s.physicAndColor(ch.hitTarget)
		switch ht {
		case HitTypeNone:
			continue
		case HitTypeElement:
			if !math.IsIntersecting(&ch.physic, physic) {
				continue
			}
			if math.IsEqualVector3(ch.ptr.Color, color) {
				ch.ptr.Physic.Mass += s.world.PlayerParameter.Penalty
				ch.ptr.Physic.Radius = math.RadiusFromVolume(ch.ptr.Physic.Mass)
			} else {
				ch.ptr.Physic.Mass += 1.0
				ch.ptr.Physic.Radius = math.RadiusFromVolume(ch.ptr.Physic.Mass)
				for i, element := range s.world.Elements {
					if element.Name == ch.hitTarget {
						s.world.Elements = append(
							s.world.Elements[:i], s.world.Elements[i+1:]...)
					}
				}
			}
		case HitTypeCharacter:
			if !math.IsIntersecting(&ch.physic, physic) {
				continue
			}
			if math.IsEqualVector3(ch.ptr.Color, color) {
				totalMass := ch.ptr.Physic.Mass + physic.Mass
				ch.ptr.Physic.Mass = totalMass / 2
				ch.ptr.Physic.Radius = math.RadiusFromVolume(totalMass / 2)
				physic.Mass = totalMass / 2
				physic.Radius = math.RadiusFromVolume(totalMass / 2)
			} else {
				if ch.ptr.Physic.Mass < physic.Mass {
					ch.ptr.Physic.Mass -= s.world.PlayerParameter.EatSpeed
					ch.ptr.Physic.Radius = math.RadiusFromVolume(ch.ptr.Physic.Mass)
					physic.Mass += s.world.PlayerParameter.EatSpeed
					physic.Radius = math.RadiusFromVolume(physic.Mass)
				} else {
					ch.ptr.Physic.Mass += s.world.PlayerParameter.EatSpeed
					ch.ptr.Physic.Radius = math.RadiusFromVolume(ch.ptr.Physic.Mass)
					physic.Mass -= s.world.PlayerParameter.EatSpeed
					physic.Radius = math.RadiusFromVolume(physic.Mass)
				}
			}
		}
	}
}

func (s *DarwinService) Ping(
	ctx context.Context, pingRequest *proto.PingRequest) (
	*proto.PingResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return &proto.PingResponse{
		Value:           pingRequest.GetValue(),
		Time:            math.TimeSecondNow(),
		PlayerParameter: s.world.PlayerParameter,
	}, nil
}

func (s *DarwinService) ReportInGame(
	ctx context.Context, reportRequest *proto.ReportInGameRequest) (
	*proto.ReportInGameResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, err := s.peer(ctx)
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
					CharacterHit{
						ptr:       character,
						physic:    *reportRequest.Physic,
						hitTarget: reportRequest.PotentialHit})
			}
			s.peerChars[p.Addr.String()] = PeerClient{character.Name, math.TimeSecondNow()}
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
	s.mu.Lock()
	defer s.mu.Unlock()
	p, err := s.peer(ctx)
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
	planet := s.planet()
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
	s.peerChars[p.Addr.String()] = PeerClient{createRequest.Name, math.TimeSecondNow()}
	s.world.Characters = append(s.world.Characters, newCharacter)
	return &proto.CreateCharacterResponse{ReturnEnum: proto.ReturnEnum_RETURN_OK}, nil
}

func (s *DarwinService) Update(
	req *proto.UpdateRequest, stream proto.DarwinService_UpdateServer) error {
	for {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("[%.3f] Processing...\n", math.TimeSecondNow())
		s.mu.Lock()
		s.removeDisconnected()
		s.checkHit()
		fmt.Printf("[%d] Elements.\n", len(s.world.Elements))
		err := stream.Send(&proto.UpdateResponse{
			Characters: s.world.Characters,
			Elements:   s.world.Elements,
			Time:       math.TimeSecondNow()})
		s.mu.Unlock()
		if err != nil {
			return err
		}
	}
}
