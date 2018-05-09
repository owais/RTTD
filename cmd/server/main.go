package main

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/owais/RTTD/pkg/api/http"
	"github.com/owais/RTTD/pkg/teams/slack"
)

func main() {

	var slackToken string
	var port string

	app := cli.NewApp()
	app.Name = "RTTD"
	app.Description = "Remote Teams Timezone Dashboard"
	app.Usage = app.Description

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "slack-api-token",
			Usage:       "Slack REST API Token",
			Destination: &slackToken,
			EnvVar:      "SLACK_API_TOKEN",
		},
		cli.StringFlag{
			Name:        "port",
			Usage:       "PORT to start web server on",
			Destination: &port,
			EnvVar:      "PORT",
			Value:       "5000",
		},
	}

	app.Action = func(c *cli.Context) error {
		if slackToken == "" {
			return errors.New("slack-api-token not provided")
		}
		endpoint := "https://slack.com/api/users.list?token=" + slackToken
		team := slack.NewTeam(endpoint)
		http.Start(team, port)
		return nil

	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
