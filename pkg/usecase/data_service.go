package usecase

import (
	"log"
	"math/rand"
	"time"

	pb "github.com/zkhrg/go_team00/pkg/api/pb"

	"github.com/google/uuid"
)

type DataService struct {
	pb.UnimplementedDataStreamServer
}

func NewDataService() *DataService {
	return &DataService{}
}

func (s *DataService) StreamData(req *pb.StreamRequest, stream pb.DataStream_StreamDataServer) error {
	sessionID := uuid.New().String()

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	expectedValue := rng.Float64()*20 - 10
	stdDev := rng.Float64()*(1.5-0.3) + 0.3

	log.Printf("Session ID: %s, Expected Value: %f, StdDev: %f", sessionID, expectedValue, stdDev)

	for {
		frequency := rng.NormFloat64()*stdDev + expectedValue
		currentTimestamp := float64(time.Now().UTC().Unix())

		if err := stream.Send(&pb.DataMessage{
			SessionId:        sessionID,
			Frequency:        frequency,
			CurrentTimestamp: currentTimestamp,
		}); err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}
}
