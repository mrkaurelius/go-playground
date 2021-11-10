package user

type User struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

var Users = []User{
	{"1", "asdf1234"},
	{"2", "gojg4567"},
}
