package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jamie-stackhouse/letsgo/Server/controllers"
)

type Config struct {
	port string
}

var (
	config     *Config
	controller *controllers.Controller = controllers.NewController()
)

func main() {
	mux := http.NewServeMux()
	controller.Attach(mux)

	config := Config{
		port: ":8080",
	}

	fmt.Printf("Listening on port: %s\n", config.port)

	if err := http.ListenAndServe(config.port, mux); err != nil {
		log.Fatal(err)
	}
}
