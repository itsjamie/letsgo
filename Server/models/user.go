package models

type User struct {
	Name string `json:"name"`
}

func NewUser(name string) *User {
	return &User{Name: name}
}
