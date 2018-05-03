package web

import (
	c "github.com/owais/rendr/pkg/components"

	"github.com/owais/RTTD/pkg/teams"
)

func userComponent(user teams.User) c.Renderer {
	return c.Div(
		c.Img(
			c.Attr("src", user.Image()),
			c.Style("border-radius", "100%"),
		),
		c.Div(c.Strong(c.Text(user.DisplayName()))),
		c.Div(c.Small(c.Text("@"+user.UserName()))),
	).Append(
		c.Styles(map[string]string{
			"padding": "15px",
			"height":  "80px",
			"width":   "80px",
		})...,
	)
}
