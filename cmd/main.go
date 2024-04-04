package main

import (
	"github.com/matisiekpl/stravactl/internal/controller"
	"github.com/matisiekpl/stravactl/internal/repository"
	"github.com/matisiekpl/stravactl/internal/service"
	strava "github.com/obalunenko/strava-api/client"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

const stravaAccessTokenEnvKey = "STRAVA_ACCESS_TOKEN"
const stravaSettingsPageUrl = "https://www.strava.com/settings/api"

func main() {
	if os.Getenv(stravaAccessTokenEnvKey) == "" {
		logrus.Fatalf("%s environment variable is not set. Go to %s and grab access token", stravaAccessTokenEnvKey, stravaSettingsPageUrl)
	}

	apiClient, err := strava.NewAPIClient(os.Getenv(stravaAccessTokenEnvKey))
	if err != nil {
		logrus.Fatalf("unable to initialize strava client: %w", err)
	}
	repositories := repository.NewRepositories(apiClient)
	services := service.NewServices(repositories)
	controllers := controller.NewControllers(services)

	app := &cli.App{
		Name:  "stravactl",
		Usage: "Strava CLI",
		Commands: []*cli.Command{
			{
				Name:    "get",
				Aliases: []string{"g", "ge"},
				Usage:   "get resource",
				Subcommands: []*cli.Command{
					{
						Name:    "activities",
						Aliases: []string{"a", "activity"},
						Usage:   "get activities",
						Action: func(ctx *cli.Context) error {
							return controllers.Activity().List(ctx.Context)
						},
					},
				},
			},
			{
				Name:    "show",
				Aliases: []string{"s", "sh"},
				Usage:   "show resource",
				Subcommands: []*cli.Command{
					{
						Name:    "activity",
						Aliases: []string{"a", "activities"},
						Usage:   "show activity",
						Args:    true,
						Action: func(ctx *cli.Context) error {
							return controllers.Activity().Show(ctx.Context, ctx.Args().Get(0))
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
