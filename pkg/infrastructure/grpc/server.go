package grpc

import (
	"log"
	"net"

	pb "github.com/zkhrg/go_team00/pkg/api/pb"
	"github.com/zkhrg/go_team00/pkg/config"
	"github.com/zkhrg/go_team00/pkg/usecase"

	"google.golang.org/grpc"
)

func RunGRPCServer(cfg config.Config, dataService *usecase.DataService) {
	lis, err := net.Listen(cfg.GRPCProt, cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterDataStreamServer(grpcServer, dataService)

	log.Printf("gRPC server is running on port %s", cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
