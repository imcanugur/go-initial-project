package repository

import (
	"go-initial-project/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	*BaseRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository[entity.User](db),
	}
}

func (ur *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
