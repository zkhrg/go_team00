package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/zkhrg/go_team00/pkg/api/pb"
	"github.com/zkhrg/go_team00/pkg/config"
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
	cfg := config.NewConfig()
	flag.Float64Var(&anomalyCoefficient, "k", 0, "Standard deviation anomaly coefficient")
	flag.Parse()

	if anomalyCoefficient <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	conn, err := grpc.NewClient(
		cfg.GRPCAddr+cfg.GRPCPort,
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

	go Detecting(cfg, freqChan, anomalyCoefficient, msg.GetSessionId()) // нужно бы контекст пробрасывать для Graceful Shutdown

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

func Detecting(cfg config.Config, freqChan chan float64, anomalyCoefficient float64, sessionID string) {
	db, err := gorm.Open(postgres.Open(cfg.PGConn), &gorm.Config{})
	for i := 0; err != nil; i++ {
		time.Sleep(1 * time.Second)
		db, err = gorm.Open(postgres.Open(cfg.PGConn), &gorm.Config{})
		if i > cfg.RetryCount {
			log.Fatal("Ошибка подключения к базе данных:", err)
		}
	}
	rep := database.NewGormAnomalyRepository(db)
	anomServ := usecase.NewAnomalyService(rep)
	anomaliesDetector := usecase.NewAnomaliesDetector(anomalyCoefficient, freqCountToCalculateParameters, anomServ, sessionID)
	anomaliesDetector.DetectAnomalies(freqChan)
}
