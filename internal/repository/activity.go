package repository

import (
	"context"
	"github.com/matisiekpl/stravactl/internal/model"
	strava "github.com/obalunenko/strava-api/client"
	"time"
)

type ActivityRepository interface {
	List(ctx context.Context) ([]model.Activity, error)
}

type activityRepository struct {
	client *strava.APIClient
}

func newActivityRepository(client *strava.APIClient) ActivityRepository {
	return &activityRepository{client}
}
func (a *activityRepository) List(ctx context.Context) ([]model.Activity, error) {
	response, err := a.client.Activities.GetLoggedInAthleteActivities(ctx)
	if err != nil {
		return nil, err
	}
	var activities []model.Activity

	for _, item := range response {
		activities = append(activities, model.Activity{
			ID:                 item.ID,
			AverageHeartRate:   item.HeartRateDetails.AverageHeartrate,
			MaxHeartRate:       item.HeartRateDetails.MaxHeartrate,
			HasHeartRate:       item.HeartRateDetails.HasHeartrate,
			AverageSpeed:       item.AverageSpeed,
			AverageWatts:       item.AverageWatts,
			Distance:           item.Distance,
			ElapsedTime:        item.ElapsedTime,
			MovingTime:         item.MovingTime,
			Name:               item.Name,
			StartDateLocal:     time.Time(item.StartDateLocal),
			TotalElevationGain: item.TotalElevationGain,
			SportType:          model.SportType(item.SportType),
		})
	}
	return activities, nil
}
