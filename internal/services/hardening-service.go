package services

import (
	"fmt"
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
	fmt.Println("Getting scores")
	lines := hardening.GetLynisScores()
	fmt.Println("Socres obtained")
	fmt.Println(lines)
	result := &grpc.LynisScore{
		Score: "",
		Log: lines,
	}
	return result, nil
}
