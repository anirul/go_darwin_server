package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	darwin "github.com/anirul/go_darwin_server/darwin"
	server "github.com/anirul/go_darwin_server/server"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost: %d", *port))
	if err != nil {
		log.Fatalf("failed to listen to: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	darwin.RegisterDarwinServiceServer(grpcServer, server.NewDarwinService())
	grpcServer.Serve(lis)
}
