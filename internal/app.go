package internal

import (
	"fmt"
	"github.com/antonioalfa22/egida-api-worker/internal/server"
	pb "github.com/antonioalfa22/egida-api-worker/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Run() {
	fmt.Println("Go gRPC Running on port 8128")
	fmt.Println("==================>")

	// Host grpc service
	listen, err := net.Listen("tcp", ":8128")
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	// gRPC server
	s := grpc.NewServer()
	pb.RegisterServicesServer(s, server.NewServicesServer())
	pb.RegisterPackagesServer(s, server.NewPackagesServer())
	pb.RegisterHardeningScoresServer(s, server.NewHardeningServer())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}