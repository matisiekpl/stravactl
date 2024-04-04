package controller

import "github.com/matisiekpl/stravactl/internal/service"

type Controllers interface {
	Activity() ActivityController
}

type controllers struct {
	activity ActivityController
}

func NewControllers(services service.Services) Controllers {
	activity := newActivityController(services.Activity())
	return &controllers{
		activity: activity,
	}
}

func (c controllers) Activity() ActivityController {
	return c.activity
}
