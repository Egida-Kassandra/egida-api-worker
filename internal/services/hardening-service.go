package services

import grpc "github.com/antonioalfa22/egida-api-worker/proto"

type HardeningService struct {}
var hardeningService *HardeningService

func GetHardeningService() *HardeningService {
	if hardeningService == nil {
		hardeningService = &HardeningService{}
	}
	return hardeningService
}

func (s HardeningService) GetLynisScore() (*grpc.LynisScore, error) {
	return nil, nil
}
