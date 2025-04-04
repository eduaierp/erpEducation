package server

import (
	"log"
	"net"

	pb "hr-service/proto"

	"google.golang.org/grpc"
)

func StartGRPCServer(h pb.HRServiceServer) {
	lis, err := net.Listen("tcp", ":50056")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHRServiceServer(grpcServer, h)

	log.Println("✅ HR Service running on port 50056...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
