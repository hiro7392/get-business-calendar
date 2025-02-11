package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/hiro7392/get-business-calendar/internal/handler"                // サービス実装があるパッケージ
	pb "github.com/hiro7392/get-business-calendar/internal/proto/business_days" // proto 定義のパッケージ
)

func main() {
	// gRPC サーバーを起動するポート番号を指定
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPC サーバーの作成
	grpcServer := grpc.NewServer()

	// handler パッケージ内の実装で BusinessDays サービスのサーバーインスタンスを生成
	businessDaysServer := handler.NewBusinessDaysServer()

	// 生成された proto のコードを利用して、gRPC サーバーにサービスを登録
	pb.RegisterBusinessDaysServer(grpcServer, businessDaysServer)

	log.Println("Server is running on port :50051")
	// サーバーの起動
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}