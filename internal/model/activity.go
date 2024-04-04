package model

import (
	"time"
)

type Activity struct {
	ID                 int64
	AverageHeartRate   float32
	MaxHeartRate       float32
	HasHeartRate       bool
	AverageSpeed       float32
	AverageWatts       float32
	Distance           float32
	ElapsedTime        int64
	MovingTime         int64
	Name               string
	StartDateLocal     time.Time
	TotalElevationGain float32
	SportType          SportType
}
