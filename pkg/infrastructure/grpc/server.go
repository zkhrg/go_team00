package grpc

import (
	"log"
	"net"

	pb "github.com/zkhrg/go_team00/pkg/api/pb"
	"github.com/zkhrg/go_team00/pkg/usecase"

	"google.golang.org/grpc"
)

func RunGRPCServer(dataService *usecase.DataService) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterDataStreamServer(grpcServer, dataService)

	log.Println("gRPC server is running on port :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
