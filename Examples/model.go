package main

import (
	"encoding/json"
	"fmt"

	"github.com/jamie-stackhouse/letsgo/Server/models"
)

//PICK1 OMIT
func main() {
	user := models.NewUser("Jeff")
	userData, _ := json.Marshal(user)
	fmt.Printf("%s\n", userData)
}

//END1 OMIT
