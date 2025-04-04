package server

import (
	"log"
	"net"

	pb "communication-service/proto"

	"google.golang.org/grpc"
)

func StartGRPCServer(h pb.CommunicationServiceServer) {
	lis, err := net.Listen("tcp", ":50059")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCommunicationServiceServer(grpcServer, h)

	log.Println("✅ Communication Service running on port 50059...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
