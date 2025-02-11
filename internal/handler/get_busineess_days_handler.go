package handler

import (
	"context"
	pb "github.com/hiro7392/get-business-calendar/internal/proto/generarted"
)

type businessDaysServer struct {
	pb.UnimplementedBusinessDaysServer
}

func NewBusinessDaysServer() pb.ServiceRegistrar {
	return &businessDaysServer{}
}

func (s *businessDaysServer) GetBusinessDays(ctx context.Context, req *pb.BusinessDaysRequest) (*pb.BusinessDaysResponse, error) {
	// 固定の営業日データを作成 
	days := []*pb.BusinessDay{
		{Date: "2025-02-03", Weekday: 1}, // 月曜日
		{Date: "2025-02-04", Weekday: 2}, // 火曜日
		{Date: "2025-02-05", Weekday: 3}, // 水曜日
		{Date: "2025-02-06", Weekday: 4}, // 木曜日
		{Date: "2025-02-07", Weekday: 5}, // 金曜日
		{Date: "2025-02-10", Weekday: 1}, // 月曜日
		{Date: "2025-02-11", Weekday: 2}, // 火曜日
		{Date: "2025-02-12", Weekday: 3}, // 水曜日
		{Date: "2025-02-13", Weekday: 4}, // 木曜日
		{Date: "2025-02-14", Weekday: 5}, // 金曜日
		{Date: "2025-02-17", Weekday: 1}, // 月曜日
		{Date: "2025-02-18", Weekday: 2}, // 火曜日
		{Date: "2025-02-19", Weekday: 3}, // 水曜日
		{Date: "2025-02-20", Weekday: 4}, // 木曜日
		{Date: "2025-02-21", Weekday: 5}, // 金曜日
		{Date: "2025-02-24", Weekday: 1}, // 月曜日
		{Date: "2025-02-25", Weekday: 2}, // 火曜日
		{Date: "2025-02-26", Weekday: 3}, // 水曜日
		{Date: "2025-02-27", Weekday: 4}, // 木曜日
		{Date: "2025-02-28", Weekday: 5}, // 金曜日
	}

	// 固定のレスポンスを作成して返す
	return &pb.BusinessDaysResponse{
		BusinessDays: days,
	}, nil
}