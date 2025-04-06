package server

import (
	"log"
	"net"

	pb "student-service/proto"

	"google.golang.org/grpc"
)

func StartGRPCServer(h pb.StudentServiceServer) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStudentServiceServer(s, h)

	log.Println("Student Service running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
