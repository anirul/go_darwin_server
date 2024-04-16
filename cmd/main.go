package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	darwin "github.com/anirul/go_darwin_server/darwin"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type darwinServer struct {
	darwin.UnimplementedDarwinServiceServer
}

func newServer() *darwinServer {
	s := &darwinServer{}
	return s
}

func main() {
	flag.Parse()
	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost: %d", *port))
	if err != nil {
		log.Fatalf("failed to listen to: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	darwin.RegisterDarwinServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
