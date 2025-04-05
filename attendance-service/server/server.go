package server

import (
	"log"
	"net"

	pb "attendance-service/proto"

	"google.golang.org/grpc"
)

func StartGRPCServer(h pb.AttendanceServiceServer) {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAttendanceServiceServer(grpcServer, h)

	log.Println("✅ Attendance Service running on port 50054...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
