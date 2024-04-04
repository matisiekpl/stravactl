package service

import "github.com/matisiekpl/stravactl/internal/repository"

type Services interface {
	Activity() ActivityService
}

type services struct {
	activity ActivityService
}

func NewServices(repositories repository.Repositories) Services {
	activity := newActivityService(repositories.Activities())
	return &services{
		activity: activity,
	}
}

func (s services) Activity() ActivityService {
	return s.activity
}
