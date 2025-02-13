package handler

import (
	"context"

	pb "github.com/hiro7392/get-business-calendar/internal/proto/health_check"
)


type healthCheckServer struct {
	pb.UnimplementedHealthServer
}

func NewHealthCheckServer() pb.healthCheckServer {
	return &healthCheckServer{}
}

func (s *healthCheckServer) GetHealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse,error){
	return &pb.HealthCheckResponse{
	},nil
}