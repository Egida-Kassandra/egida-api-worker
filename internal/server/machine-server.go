package server

import (
	"context"
	"github.com/antonioalfa22/egida-api-worker/internal/services"
	grpc "github.com/antonioalfa22/egida-api-worker/proto"
)

type machineinfoServer struct {}

func NewMachineInfoServer() grpc.MachineInfoServer{
	return &machineinfoServer{}
}

func (m machineinfoServer) ListOpenPorts(ctx context.Context, req *grpc.ListOpenPortsReq) (*grpc.ListPort, error) {
	s := services.GetMachineInfoService()
	return s.GetListenOpenPorts()
}

func (m machineinfoServer) MachineHasGUI(ctx context.Context, req *grpc.HasGUIReq) (*grpc.HasGUI, error) {
	s := services.GetMachineInfoService()
	return s.MachineHasGUI()
}
