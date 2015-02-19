package main

import (
	"log"
	"net/http"
)

//PICK1 OMIT
func main() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

//END1 OMIT
