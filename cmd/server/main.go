package main

import (
	"log"

	"github.com/zkhrg/go_team00/pkg/config"
	"github.com/zkhrg/go_team00/pkg/infrastructure/grpc"
	"github.com/zkhrg/go_team00/pkg/usecase"
)

func main() {
	log.Println("Starting gRPC server...")

	dataService := usecase.NewDataService()
	cfg := config.NewConfig()
	grpc.RunGRPCServer(cfg, dataService)
}
