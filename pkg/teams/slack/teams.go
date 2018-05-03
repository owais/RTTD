package slack

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/owais/RTTD/pkg/teams"
)

type Team struct {
	apiEndpoint string
	timezones   []teams.Timezone
}

func NewTeam(apiEndpoint string) teams.Team {
	return &Team{apiEndpoint: apiEndpoint}
}

func (t *Team) Timezones() []teams.Timezone {
	return t.timezones
}

func (t *Team) Fetch() ([]byte, error) {
	resp, err := http.Get(t.apiEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (t *Team) Refresh() error {
	// response, err := ioutil.ReadFile("users.json")
	response, err := t.Fetch()

	if err != nil {
		return err
	}

	r := struct {
		OK      bool         `json:"ok"`
		Members []userRecord `json:"members"`
	}{}

	json.Unmarshal(response, &r)

	if r.OK == false {
		return errors.New("Failed to load data from slack")
	}

	zones := map[int]*Timezone{}
	for _, user := range r.Members {
		if user.Deleted == false {
			zone, ok := zones[user.TzOffset]
			if !ok {
				zone = &Timezone{
					code:   user.Tz,
					label:  user.TzLabel,
					offset: user.TzOffset,
					users:  []teams.User{},
				}
				zones[user.TzOffset] = zone
			}
			zone.users = append(zone.users, &User{user})
		}
	}

	for _, zone := range zones {
		t.timezones = append(t.timezones, zone)
	}
	sort.Slice(t.timezones, func(i, j int) bool {
		return t.timezones[i].Offset() < t.timezones[j].Offset()
	})
	return nil
}
