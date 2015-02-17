package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jamie-stackhouse/letsgo/Server/controllers"
)

func main() {
	var (
		host = flag.String("HOST", "localhost", "The host name the server is running on")
		port = flag.String("PORT", ":8080", "The server port")
	)

	config := controllers.ControllerConfig{
		Host: *host,
		Port: *port,
	}

	controller := controllers.NewController(&config)
	mux := http.NewServeMux()

	controller.Attach(mux)

	fmt.Printf("Listening on http://%s%s\n", *host, *port)

	if err := http.ListenAndServe(*port, mux); err != nil {
		log.Fatal(err)
	}
}
