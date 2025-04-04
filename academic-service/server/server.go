package server

import (
    "log"
    "net"
    "google.golang.org/grpc"
    pb "academic-service/proto"
    "academic-service/handler"
)

func StartGRPCServer(h pb.AcademicServiceServer) {
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterAcademicServiceServer(s, h)

    log.Println("Academic Service running on port 50052")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
