package main

import (
	"log"
	"net"

	"github.com/hiro7392/get-business-calendar/internal/handler"                // サービス実装があるパッケージ
	pb "github.com/hiro7392/get-business-calendar/internal/proto/business_days" // proto 定義のパッケージ
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// handler パッケージ内の実装で BusinessDays サービスのサーバーインスタンスを生成
	businessDaysServer := handler.NewBusinessDaysServer()

	// 生成された proto のコードを利用して、gRPC サーバーにサービスを登録
	pb.RegisterBusinessDaysServer(grpcServer, businessDaysServer)

	reflection.Register(grpcServer)

	log.Println("Server is running on port :9090")
	// サーバーの起動
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}