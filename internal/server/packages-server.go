package server

import (
	"context"
	"github.com/antonioalfa22/egida-api-worker/internal/services"
	grpc "github.com/antonioalfa22/egida-api-worker/proto"
)

type packagesServer struct {}

func NewPackagesServer() grpc.PackagesServer{
	return &packagesServer{}
}

func (p packagesServer) ListAllPackages(ctx context.Context, req *grpc.ListPackagesReq) (*grpc.ListPackage, error) {
	s := services.GetPackagesService()
	return s.GetAllPackages()
}