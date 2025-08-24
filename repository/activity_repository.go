package repository

import (
	"go-initial-project/entity"

	"gorm.io/gorm"
)

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

func (r *ActivityRepository) Create(activity *entity.Activity) error {
	return r.db.Create(activity).Error
}
