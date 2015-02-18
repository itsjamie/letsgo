package models

import (
	"fmt"
)

func ExampleNewUser() {
	user := NewUser("jeff")

	fmt.Printf("%s", user.Name)
	// Output: jeff
}
