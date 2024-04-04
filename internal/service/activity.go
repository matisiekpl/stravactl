package service

import (
	"context"
	"github.com/matisiekpl/stravactl/internal/model"
	"github.com/matisiekpl/stravactl/internal/repository"
)

type ActivityService interface {
	List(ctx context.Context) ([]model.Activity, error)
}

type activityService struct {
	activityRepository repository.ActivityRepository
}

func newActivityService(activityRepository repository.ActivityRepository) ActivityService {
	return &activityService{activityRepository}
}

func (a activityService) List(ctx context.Context) ([]model.Activity, error) {
	return a.activityRepository.List(ctx)
}
