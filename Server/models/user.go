package models

// User stores a users name
type User struct {
	Name string `json:"name"`
}

// NewUser accepts a users name and returns a new User type
func NewUser(name string) *User {
	return &User{Name: name}
}
