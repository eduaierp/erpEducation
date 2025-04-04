package server

import (
	"log"
	"net"

	pb "exam-service/proto"

	"google.golang.org/grpc"
)

func StartGRPCServer(h pb.ExamServiceServer) {
	lis, err := net.Listen("tcp", ":50057")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterExamServiceServer(grpcServer, h)

	log.Println("✅ Exam Service running on port 50057...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
