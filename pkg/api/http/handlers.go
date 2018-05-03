package http

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/phogolabs/parcello"

	"github.com/owais/RTTD/pkg/teams"
	"github.com/owais/RTTD/pkg/ui/web"

	_ "github.com/owais/RTTD/static" // bundle assets
)

var port string
var slackToken string

type handler func(t teams.Team, w http.ResponseWriter, r *http.Request)

func withTeams(team teams.Team, f handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(team, w, r)
	}
}

func indexHandler(t teams.Team, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	file, err := parcello.Open("index.html")
	if err != nil {
		http.Error(w, "Error finding template", 500)
		return
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading template", 500)
		return
	}

	err = t.Refresh()
	if err != nil {
		http.Error(w, "error fetch teams", 500)
		return
	}

	app := web.App{Team: t}
	tpl, err := template.New("index").Parse(string(contents))
	if err != nil {
		http.Error(w, "Error parsing template", 500)
		return
	}

	tpl.Execute(w, template.HTML(app.RenderToText()))
}

func fetchFromSlack(t teams.Team, w http.ResponseWriter, r *http.Request) {
	contents, err := t.Fetch()
	if err != nil {
		http.Error(w, "Failed to load data from Slack", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(contents)
}

func Start(team teams.Team) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(parcello.Root("/"))))

	http.HandleFunc("/api/slack/fetch/", withTeams(team, fetchFromSlack))
	http.HandleFunc("/", withTeams(team, indexHandler))

	port = "5000"
	fmt.Println("Starting server on port " + port)
	http.ListenAndServe(":"+port, nil)
}
