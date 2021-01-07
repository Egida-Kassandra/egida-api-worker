package server

import (
	"context"
	"github.com/antonioalfa22/egida-api-worker/internal/services"
	grpc "github.com/antonioalfa22/egida-api-worker/proto"
)

type servicesServer struct {}

func NewServicesServer() grpc.ServicesServer{
	return &servicesServer{}
}

func (s servicesServer) ListAllServices(ctx context.Context, req *grpc.ListServicesReq) (*grpc.ListService, error) {
	ser := services.GetServicesService()
	return ser.GetAllServices()
}

func (s servicesServer) ListRunningServices(ctx context.Context, req *grpc.ListServicesReq) (*grpc.ListService, error) {
	ser := services.GetServicesService()
	return ser.GetRunningServices()
}

func (s servicesServer) ListStoppedServices(ctx context.Context, req *grpc.ListServicesReq) (*grpc.ListService, error) {
	ser := services.GetServicesService()
	return ser.GetStoppedServices()
}

func (s servicesServer) GetService(ctx context.Context, req *grpc.ServiceNameReq) (*grpc.Service, error) {
	ser := services.GetServicesService()
	return ser.GetService(req)
}