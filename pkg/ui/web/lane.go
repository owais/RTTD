package web

import (
	"time"

	c "github.com/owais/rendr/pkg/components"

	"github.com/owais/RTTD/pkg/teams"
)

var headerStyle = map[string]string{
	"background": "rgba(0,0,0,.5)",
	"color":      "#fff",
	"padding":    "10px",
	"font-size":  "115%",
}

// func laneComponent(offset int, users []models.User) c.Renderer {
func lane(tz teams.Timezone) c.Renderer {
	users := tz.Users()
	now := time.Now().UTC()
	local := now.Add(time.Duration(tz.Offset()) * time.Second)

	styles := map[string]string{
		"background": "#fff",
		"color":      "#000",
	}
	if local.Hour() < 8 {
		styles["background"] = "#222c3b"
		styles["color"] = "#fff"
	}

	header := c.Div(
		c.Div(c.Text(local.Format("Monday"))),
		c.Div(c.Strong(c.Text(local.Format("3:04pm")))),
	).Append(c.Styles(headerStyle)...)

	children := c.Styles(styles)
	children = append(children, header)
	for _, user := range users {
		children = append(children, userComponent(user))
	}

	return c.Div(children...)
}
