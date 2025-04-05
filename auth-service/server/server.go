package server

import (
	"auth-service/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartGRPCServer(handler proto.AuthServiceServer) {
	listener, err := net.Listen("tcp", ":50061")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterAuthServiceServer(grpcServer, handler)

	log.Println("✅ Auth Service is running on port 50061...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
