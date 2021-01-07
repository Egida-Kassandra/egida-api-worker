package services

import (
	"github.com/antonioalfa22/egida-api-worker/pkg/packages"
	grpc "github.com/antonioalfa22/egida-api-worker/proto"
	"github.com/antonioalfa22/go-utils/collections"
	"strings"
)

type PackagesService struct {}
var packagesService *PackagesService

func GetPackagesService() *PackagesService {
	if packagesService == nil {
		packagesService = &PackagesService{}
	}
	return packagesService
}

func (s PackagesService) GetAllPackages() (*grpc.ListPackage, error) {
	lines := packages.GetAllPackages()
	packs := collections.Map(collections.Filter(lines, func(x string) bool {
		return x != "" && x != "Listing..." && x != "WARNING: apt does not have a stable CLI interface. Use with caution in scripts."
	}).([]string), func(x string) *grpc.Package {
		return &grpc.Package{Name: strings.Split(x, "/")[0]}
	}).([]*grpc.Package)
	result := &grpc.ListPackage{
		Packages: packs,
	}
	return result, nil
}
