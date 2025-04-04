package server

import (
	"log"
	"net"

	pb "auth-service/proto"

	"google.golang.org/grpc"
)

func StartGRPCServer(h pb.AuthServiceServer) {
	lis, err := net.Listen("tcp", ":50061")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, h)

	log.Println("✅ Auth Service running on port 50061...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
