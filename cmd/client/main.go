package main

import (
	"context"
	"log"

	pb "github.com/zkhrg/go_team00/pkg/api/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataStreamClient(conn)

	ctx := context.Background()
	resp, err := client.StreamData(ctx, &pb.StreamRequest{})
	if err != nil {
		log.Fatalf("Error while calling StreamData: %v", err)
	}
	log.Printf("Response: %v", resp)
}
