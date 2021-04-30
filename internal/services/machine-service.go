package services

import (
	"github.com/antonioalfa22/egida-api-worker/pkg/machine"
	grpc "github.com/antonioalfa22/egida-api-worker/proto"
	"github.com/antonioalfa22/go-utils/collections"
)

type MachineInfoService struct {}
var machineinfoService *MachineInfoService

func GetMachineInfoService() *MachineInfoService {
	if machineinfoService == nil {
		machineinfoService = &MachineInfoService{}
	}
	return machineinfoService
}

func (s MachineInfoService) GetListenOpenPorts() (*grpc.ListPort, error) {
	openports := machine.ListOpenPorts()
	ports := collections.Map(openports, func(x int64) *grpc.Port {
		return &grpc.Port{Number: x}
	}).([]*grpc.Port)
	result := &grpc.ListPort{
		Ports: ports,
	}
	return result, nil
}

func (s MachineInfoService) MachineHasGUI() (*grpc.HasGUI, error) {
	hasgui := machine.MachineHasGUI()
	result := &grpc.HasGUI{
		Gui: hasgui,
	}
	return result, nil
}