package repository

import "github.com/zkhrg/go_team00/pkg/domain/model"

// AnomalyRepository — это интерфейс, который определяет методы для работы с аномалиями.
type AnomalyRepository interface {
	Save(anomaly *model.Anomaly) error
	GetByID(id uint) (*model.Anomaly, error)
}
