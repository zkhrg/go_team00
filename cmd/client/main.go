package main

import (
	"context"
	"fmt"
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

	msg, err := resp.Recv()
	if err != nil {
		log.Fatalf("Error while calling Recv: %v", err)
	}
	fmt.Printf(
		"received first message | session:  %s | timestamp: %f | frequency: %f\n",
		msg.GetSessionId(),
		msg.GetCurrentTimestamp(),
		msg.GetFrequency(),
	)

	for {
		msg, err = resp.Recv()
		if err != nil {
			log.Fatalf("Error while receiving msg from server")
		}
		fmt.Printf(
			"timestamp: %f | frequency: %f\n",
			msg.GetCurrentTimestamp(),
			msg.GetFrequency(),
		)
	}
}
