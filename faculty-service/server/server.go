package server

import (
    "log"
    "net"

    "google.golang.org/grpc"
    pb "faculty-service/proto"
)

func StartGRPCServer(h pb.FacultyServiceServer) {
    lis, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterFacultyServiceServer(grpcServer, h)

    log.Println("✅ Faculty Service is running on port 50053...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
