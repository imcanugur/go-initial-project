package service

import (
	"go-initial-project/entity"
	"go-initial-project/repository"
)

type UserService struct {
	*BaseService[entity.User]
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		BaseService: &BaseService[entity.User]{repo: repo},
	}
}

func (us *UserService) FindByEmail(email string) (*entity.User, error) {
	repo := us.BaseService.repo.(*repository.UserRepository)
	return repo.FindByEmail(email)
}
