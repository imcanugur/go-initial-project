package service

import (
	"go-initial-project/entity"
	"go-initial-project/repository"
)

type ActivityService struct {
	repo *repository.ActivityRepository
}

func NewActivityService(repo *repository.ActivityRepository) *ActivityService {
	return &ActivityService{repo: repo}
}

func (s *ActivityService) Log(activity *entity.Activity) error {
	return s.repo.Create(activity)
}
