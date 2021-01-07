package services

import (
	"github.com/antonioalfa22/egida-api-worker/pkg/services"
	grpc "github.com/antonioalfa22/egida-api-worker/proto"
	"github.com/antonioalfa22/go-utils/collections"
	"strings"
)

type ServiceService struct {}
var serviceService *ServiceService

func GetServicesService() *ServiceService {
	if serviceService == nil {
		serviceService = &ServiceService{}
	}
	return serviceService
}

func (s ServiceService) GetAllServices() (*grpc.ListService, error) {
	lines := services.GetAllServices()
	servs := collections.Map(collections.Filter(lines, func(x string) bool {
		return x != ""
	}).([]string), func(x string) *grpc.Service {
		var status string
		if strings.Split(x, "  ")[0] == " [ + ]" { status = "Running"} else { status = "Stopped"}
		return &grpc.Service{
			Name: strings.Split(x, "  ")[1],
			Status: status,
		}
	}).([]*grpc.Service)
	result := &grpc.ListService{
		Services: servs,
	}
	return result, nil
}

func (s ServiceService) GetRunningServices() (*grpc.ListService, error) {
	lines := services.GetAllServices()
	servs := collections.Map(collections.Filter(lines, func(x string) bool {
		return strings.Split(x, "  ")[0] == " [ + ]"
	}).([]string), func(x string) *grpc.Service {
		return &grpc.Service{
			Name: strings.Split(x, "  ")[1],
			Status: "Running",
		}
	}).([]*grpc.Service)
	result := &grpc.ListService{
		Services: servs,
	}
	return result, nil
}

func (s ServiceService) GetStoppedServices() (*grpc.ListService, error) {
	lines := services.GetAllServices()
	servs := collections.Map(collections.Filter(lines, func(x string) bool {
		return x != "" && strings.Split(x, "  ")[0] != " [ + ]"
	}).([]string), func(x string) *grpc.Service {
		return &grpc.Service{
			Name: strings.Split(x, "  ")[1],
			Status: "Stopped",
		}
	}).([]*grpc.Service)
	result := &grpc.ListService{
		Services: servs,
	}
	return result, nil
}

func (s ServiceService) GetService(req *grpc.ServiceNameReq) (*grpc.Service, error) {
	lines := services.GetAllServices()
	serv := collections.Find(collections.Filter(lines, func(x string) bool {return x != ""}).([]string),
		func(x string) bool {
			return strings.Split(x, "  ")[1] == req.Name
		}).(string)
	var status string
	if strings.Split(serv, "  ")[0] == " [ + ]" { status = "Running"} else { status = "Stopped"}
	return &grpc.Service{
		Name: strings.Split(serv, "  ")[1],
		Status: status,
	}, nil
}
