package usecase

import (
	"log"
	"time"

	"github.com/zkhrg/go_team00/pkg/domain/model"
	"github.com/zkhrg/go_team00/pkg/domain/repository"
)

type AnomalyService struct {
	repo repository.AnomalyRepository
}

func NewAnomalyService(repo repository.AnomalyRepository) *AnomalyService {
	return &AnomalyService{repo: repo}
}

func (s *AnomalyService) StoreAnomaly(sessionID string, frequency float64) error {
	anomaly := &model.Anomaly{
		SessionID: sessionID,
		Frequency: frequency,
		Timestamp: time.Now(),
	}
	err := s.repo.Save(anomaly)
	if err != nil {
		return err
	}

	log.Println("Anomaly stored:", anomaly)
	return nil
}
