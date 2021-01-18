package server

import (
	"context"
	"fmt"
	"github.com/antonioalfa22/egida-api-worker/internal/services"
	grpc "github.com/antonioalfa22/egida-api-worker/proto"
)

type hardeningServer struct {}

func NewHardeningServer() grpc.HardeningScoresServer{
	return &hardeningServer{}
}

func (h hardeningServer) GetLynisScore(ctx context.Context, req *grpc.ScoreReq) (*grpc.LynisScore, error) {
	fmt.Println("Starting service")
	s := services.GetHardeningService()
	fmt.Println("Service created")
	return s.GetLynisScore()
}