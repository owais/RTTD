package slack

type userRecord struct {
	ID       string `json:"id"`
	Deleted  bool   `json:"deleted"`
	Name     string `json:"real_name"`
	Handle   string `json:"name"`
	Tz       string `json:"tz"`
	TzLabel  string `json:"tz_label"`
	TzOffset int    `json:"tz_offset"`
	Profile  struct {
		Avatar string `json:"image_48"`
	} `json:"profile"`
}

// User represents a slack user
type User struct {
	record userRecord
}

func (u *User) ID() string {
	return u.record.ID
}

func (u *User) DisplayName() string {
	return u.record.Name
}

func (u *User) UserName() string {
	return u.record.Handle
}

func (u *User) Image() string {
	return u.record.Profile.Avatar
}
