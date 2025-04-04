package server

import (
    "log"
    "net"

    "google.golang.org/grpc"
    pb "finance-service/proto"
)

func StartGRPCServer(h pb.FinanceServiceServer) {
    lis, err := net.Listen("tcp", ":50055")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterFinanceServiceServer(grpcServer, h)

    log.Println("✅ Finance Service running on port 50055...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
