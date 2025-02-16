package storage

import (
	"gorm.io/gorm"
	"log"
)

type PredictionRepository struct {
	db *gorm.DB
}

func NewPredictionRepository(db *gorm.DB) *PredictionRepository {
	return &PredictionRepository{db: db}
}

func (r *PredictionRepository) SavePrediction(prediction *Prediction) error {
	log.Println(prediction)
	return r.db.Create(prediction).Error
}
