package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	pb "github.com/zkhrg/go_team00/pkg/api/pb"
	"github.com/zkhrg/go_team00/pkg/database"
	"github.com/zkhrg/go_team00/pkg/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const freqCountToCalculateParameters = 100

func main() {
	var anomalyCoefficient float64
	flag.Float64Var(&anomalyCoefficient, "k", 0, "Standard deviation anomaly coefficient")
	flag.Parse()

	if anomalyCoefficient <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	conn, err := grpc.NewClient(
		"data-stream-server:50051",
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

	freqChan := make(chan float64)
	defer close(freqChan)

	go Detecting(freqChan, anomalyCoefficient, msg.GetSessionId()) // нужно бы контекст пробрасывать для Graceful Shutdown

	for {
		msg, err = resp.Recv()
		if err != nil {
			log.Fatalf("Error while receiving msg from server")
		}
		freqChan <- msg.GetFrequency()
		fmt.Printf(
			"timestamp: %f | frequency: %f\n",
			msg.GetCurrentTimestamp(),
			msg.GetFrequency(),
		)
	}
}

func Detecting(freqChan chan float64, anomalyCoefficient float64, sessionID string) {
	// dsn := "host=localhost user=postgres password=postgres dbname=school21 port=5432 sslmode=disable"
	dsn := os.Getenv("POSTGRES_CONN")
	if dsn == "" {
		log.Fatal("POSTGRES_CONN environment variable is not set")
	}
	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	rep := database.NewGormAnomalyRepository(db)
	anomServ := usecase.NewAnomalyService(rep)
	anomaliesDetector := usecase.NewAnomaliesDetector(anomalyCoefficient, freqCountToCalculateParameters, anomServ, sessionID)
	anomaliesDetector.DetectAnomalies(freqChan)
}
