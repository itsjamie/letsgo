package main

import (
	"github.com/jamie-stackhouse/letsgo/Server/controllers"
	"log"
	"net/http"
)

var (
	port       string                  = ":8080"
	controller *controllers.Controller = controllers.NewController()
)

func main() {
	mux := http.NewServeMux()
	controller.Attach(mux)

	if err := http.ListenAndServe(port, mux); err != nil {

		log.Fatal("ListenAndServe: ", err)
	}
}
