package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"

	math "github.com/anirul/go_darwin_server/darwin_math"
	proto "github.com/anirul/go_darwin_server/darwin_proto"
	service "github.com/anirul/go_darwin_server/darwin_service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	playerParam = flag.String("player_parameter", "./player_parameter.json", "Player parameter json file")
)

func ParseJsonPlayerParam(f string) (*proto.WorldDatabase, error) {
	file, err := os.Open(f)
	params := &proto.WorldDatabase{}
	if err != nil {
		return params, err
	}
	defer file.Close()
	jsonData, err := io.ReadAll(file)
	if err != nil {
		return params, err
	}
	err = protojson.Unmarshal(jsonData, params)
	if err != nil {
		return params, err
	}
	return params, nil
}

func PopulateWorldDatabase(worldDatabase *proto.WorldDatabase, nbBonus int) (*proto.WorldDatabase, error) {
	planet := &proto.Element{}
	for _, e := range worldDatabase.Elements {
		if e.Name == "earth" && e.TypeEnum == proto.TypeEnum_TYPE_GROUND {
			planet = e
		}
	}
	if planet.Name == "" {
		return worldDatabase, fmt.Errorf("no planet found in the world database")
	}
	radius := math.RadiusFromVolume(1)
	radiusFromEarth := radius + planet.Physic.Radius
	for i := 0; i < nbBonus; i++ {
		randomIndex := rand.Intn(len(worldDatabase.PlayerParameter.ColorParameters))
		worldDatabase.Elements = append(worldDatabase.Elements, &proto.Element{
			Name:  fmt.Sprintf("element%d", i),
			Color: worldDatabase.PlayerParameter.ColorParameters[randomIndex].Color,
			Physic: &proto.Physic{
				Radius:   radius,
				Mass:     1,
				Position: math.Times(math.RandomNormalizeVec3(), radiusFromEarth)},
			TypeEnum: proto.TypeEnum_TYPE_UPGRADE})
	}
	return worldDatabase, nil
}

func main() {
	flag.Parse()
	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost: %d", *port))
	if err != nil {
		log.Fatalf("Failed to listen to: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	param, err := ParseJsonPlayerParam(*playerParam)
	if err != nil {
		log.Fatalf("Failed at service creation: %v", err)
	}
	world, err := PopulateWorldDatabase(param, 200)
	if err != nil {
		log.Fatalf("Failed at populate: %v", err)
	}
	proto.RegisterDarwinServiceServer(grpcServer, service.NewDarwinService(world))
	grpcServer.Serve(lis)
}
