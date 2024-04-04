package controller

import (
	"context"
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/matisiekpl/stravactl/internal/service"
	"github.com/matisiekpl/stravactl/internal/util"
	"time"
)

type ActivityController interface {
	List(ctx context.Context) error
	Show(ctx context.Context, id string) error
}

type activityController struct {
	activityService service.ActivityService
}

func newActivityController(activityService service.ActivityService) ActivityController {
	return &activityController{activityService}
}

func (a *activityController) List(ctx context.Context) error {
	activities, err := a.activityService.List(ctx)
	if err != nil {
		return err
	}
	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true
	table.AddRow("SPORT", "ID", "NAME", "TIME", "DISTANCE", "HR", "SPEED", "AGE")
	for _, activity := range activities {
		table.AddRow(
			string(activity.SportType),
			activity.ID,
			activity.Name,
			util.FormatElapsedTime(activity.ElapsedTime),
			util.FormatDistance(activity.Distance),
			util.FormatHeartRate(activity),
			util.FormatSpeed(activity),
			activity.StartDateLocal.Format(time.DateTime),
		)
	}
	fmt.Println(table)
	return nil
}

func (a *activityController) Show(ctx context.Context, id string) error {
	return util.Open("https://www.strava.com/activities/" + id)
}
