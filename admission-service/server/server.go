package server

import (
	"log"
	"net"

	pb "admission-service/proto"

	"google.golang.org/grpc"
)

func StartGRPCServer(h pb.AdmissionServiceServer) {
	lis, err := net.Listen("tcp", ":50058")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdmissionServiceServer(grpcServer, h)

	log.Println("✅ Admission Service running on port 50058...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
