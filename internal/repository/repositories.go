package repository

import (
	strava "github.com/obalunenko/strava-api/client"
)

type Repositories interface {
	Activities() ActivityRepository
}

type repositories struct {
	activity ActivityRepository
}

func NewRepositories(client *strava.APIClient) Repositories {
	activity := newActivityRepository(client)
	return &repositories{
		activity: activity,
	}
}

func (r repositories) Activities() ActivityRepository {
	return r.activity
}
