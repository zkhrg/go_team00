package database

import (
	"time"

	"github.com/zkhrg/go_team00/pkg/domain/model"
)

// AnomalyEntity — это модель базы данных с метками GORM.
type AnomalyEntity struct {
	ID        uint      `gorm:"primaryKey"`
	SessionID string    `gorm:"index"`
	Frequency float64   `gorm:"not null"`
	Timestamp time.Time `gorm:"not null"`
}

// ToDomain — преобразует GORM-модель в доменную модель.
func (e *AnomalyEntity) ToDomain() *model.Anomaly {
	return &model.Anomaly{
		ID:        e.ID,
		SessionID: e.SessionID,
		Frequency: e.Frequency,
		Timestamp: e.Timestamp,
	}
}

// FromDomain — создает GORM-модель на основе доменной модели.
func FromDomain(anomaly *model.Anomaly) *AnomalyEntity {
	return &AnomalyEntity{
		ID:        anomaly.ID,
		SessionID: anomaly.SessionID,
		Frequency: anomaly.Frequency,
		Timestamp: anomaly.Timestamp,
	}
}
