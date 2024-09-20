package database

import (
	"github.com/zkhrg/go_team00/pkg/domain/model"
	"github.com/zkhrg/go_team00/pkg/domain/repository"

	"gorm.io/gorm"
)

// GormAnomalyRepository — реализация интерфейса AnomalyRepository с использованием GORM.
type GormAnomalyRepository struct {
	db *gorm.DB
}

// NewGormAnomalyRepository создает новый репозиторий с использованием GORM.
func NewGormAnomalyRepository(db *gorm.DB) repository.AnomalyRepository {
	Migrate(db)
	return &GormAnomalyRepository{db: db}
}

// Save сохраняет аномалию в базе данных, преобразовывая её в GORM-модель.
func (r *GormAnomalyRepository) Save(anomaly *model.Anomaly) error {
	entity := FromDomain(anomaly)
	return r.db.Create(&entity).Error
}

// GetByID получает аномалию по ID и преобразует её обратно в доменную модель.
func (r *GormAnomalyRepository) GetByID(id uint) (*model.Anomaly, error) {
	var entity AnomalyEntity
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return entity.ToDomain(), nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&AnomalyEntity{})
}
