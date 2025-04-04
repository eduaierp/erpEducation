package server

import (
	"log"
	"net"

	pb "reporting-service/proto"

	"google.golang.org/grpc"
)

func StartGRPCServer(h pb.ReportingServiceServer) {
	lis, err := net.Listen("tcp", ":50060")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterReportingServiceServer(grpcServer, h)

	log.Println("✅ Reporting Service running on port 50060...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
