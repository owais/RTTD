package teams

type Timezone interface {
	Code() string
	Label() string
	Offset() int
	Users() []User
}

type User interface {
	ID() string
	DisplayName() string
	UserName() string
	Image() string
}

type Team interface {
	Timezones() []Timezone
	Fetch() ([]byte, error)
	Refresh() error
}
