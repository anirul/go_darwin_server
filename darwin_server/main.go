package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	proto "github.com/anirul/go_darwin_server/darwin_proto"
	server "github.com/anirul/go_darwin_server/server"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	playerParam = flag.String("player_parameter", "./player_parameter.json", "Player parameter json file")
)

func ParseJsonPlayerParam(f string) (*proto.PlayerParameter, error) {
	file, err := os.Open(f)
	params := &proto.PlayerParameter{}
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
	proto.RegisterDarwinServiceServer(grpcServer, server.NewDarwinService(param))
	grpcServer.Serve(lis)
}
