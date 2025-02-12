package main

import (
	"log"
	"net"
	"os"

	"github.com/hiro7392/get-business-calendar/internal/handler"                // サービス実装があるパッケージ
	pb "github.com/hiro7392/get-business-calendar/internal/proto/business_days" // proto 定義のパッケージ
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    lis, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
	
	grpcServer := grpc.NewServer()

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	businessDaysServer := handler.NewBusinessDaysServer()
	pb.RegisterBusinessDaysServer(grpcServer, businessDaysServer)

	reflection.Register(grpcServer)

	log.Printf("Server is running on port :%s",port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}