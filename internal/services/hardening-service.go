package services

import (
	"github.com/antonioalfa22/egida-api-worker/pkg/hardening"
	grpc "github.com/antonioalfa22/egida-api-worker/proto"
)

type HardeningService struct {}
var hardeningService *HardeningService

func GetHardeningService() *HardeningService {
	if hardeningService == nil {
		hardeningService = &HardeningService{}
	}
	return hardeningService
}

func (s HardeningService) GetLynisScore() (*grpc.LynisScore, error) {
	lines := hardening.GetLynisScores()
	result := &grpc.LynisScore{
		Score: "",
		Log: lines,
	}
	return result, nil
}
