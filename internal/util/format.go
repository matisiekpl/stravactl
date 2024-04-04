package util

import (
	"fmt"
	"github.com/matisiekpl/stravactl/internal/model"
	"time"
)

func FormatHeartRate(activity model.Activity) string {
	if activity.HasHeartRate {
		return fmt.Sprintf("%d", int(activity.AverageHeartRate))
	}
	return "-"
}

func FormatDistance(meters float32) string {
	if meters < 1000 {
		return fmt.Sprintf("%dm", int64(meters))
	}
	kilometers := meters / 1000
	return fmt.Sprintf("%.01fkm", kilometers)
}

func FormatElapsedTime(elapsedTime int64) string {
	minutes := int(elapsedTime / 60)
	return fmt.Sprintf("%dm", minutes)
}

func FormatSpeed(activity model.Activity) string {
	if activity.AverageSpeed == 0 {
		return "-"
	}
	switch activity.SportType {
	case model.SportTypeRun, model.SportTypeTrailRun, model.SportTypeVirtualRun:
		paceInSecondsPerKm := 1000 / activity.AverageSpeed
		paceDuration := time.Duration(paceInSecondsPerKm) * time.Second
		paceMinutes := int(paceDuration.Minutes())
		paceSeconds := int(paceDuration.Seconds()) % 60
		return fmt.Sprintf("%02d:%02d", paceMinutes, paceSeconds)
	default:
		return fmt.Sprintf("%0.1fkm/h", activity.AverageSpeed*3.6)
	}
}
