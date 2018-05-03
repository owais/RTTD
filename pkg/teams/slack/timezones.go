package slack

import "github.com/owais/RTTD/pkg/teams"

type Timezone struct {
	code   string
	label  string
	offset int
	users  []teams.User
}

func (t *Timezone) Code() string {
	return t.code
}

func (t *Timezone) Label() string {
	return t.label
}

func (t *Timezone) Offset() int {
	return t.offset
}

func (t *Timezone) Users() []teams.User {
	return t.users
}
