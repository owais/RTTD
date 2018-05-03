package web

import (
	c "github.com/owais/rendr/pkg/components"

	"github.com/owais/RTTD/pkg/teams"
)

type App struct {
	c.Component
	Team teams.Team
}

func (a *App) RenderToText() string {
	return a.Render().Text()
}

func (a App) Render() c.Renderer {
	root := a.Append(c.Styles(map[string]string{
		"display":    "flex",
		"text-align": "center",
	})...)

	for _, tz := range a.Team.Timezones() {
		root = root.Append(lane(tz))
	}
	return root
}
