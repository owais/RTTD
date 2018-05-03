package main

import (
	"fmt"
	"time"

	"github.com/owais/rendr/pkg/dom"

	"github.com/owais/RTTD/pkg/teams"
	"github.com/owais/RTTD/pkg/teams/slack"
	"github.com/owais/RTTD/pkg/ui/web"
)

func fetchPeriodically(t teams.Team, ch chan struct{}) {
	for {
		err := t.Refresh()
		if err != nil {
			fmt.Println("Error fetching users")
		} else {
			ch <- struct{}{}
		}
		time.Sleep(30 * time.Minute)
	}
}

func nextTick() time.Duration {
	return time.Duration(61 - time.Now().Second())
}

func renderLoop(t teams.Team, ch chan struct{}) {
	app := web.App{Team: t}
	for {
		select {
		case <-time.After(nextTick() * time.Second):
			dom.Render("#app", app)
		case <-ch:
			dom.Render("#app", app)
		}
	}
}

func main() {
	t := slack.NewTeam("/api/slack/fetch/")
	ch := make(chan struct{})
	go renderLoop(t, ch)
	go fetchPeriodically(t, ch)
}
