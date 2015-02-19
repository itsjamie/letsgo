package controllers

import (
	"flag"
	"fmt"
	"net/http"
)

func ExampleControllerConfig() {
	// Generate a basic config to specify the host, port, and template location
	config := ControllerConfig{
		Host:      "localhost",
		Port:      ":8080",
		Templates: "templates/",
	}

	url := fmt.Sprintf("http://%s%s", config.Host, config.Port)
	fmt.Printf(url)

	// Output: http://localhost:8080
}

func ExampleControllerConfig_flags() {
	var (
		host      = flag.String("HOST", "localhost", "The host name the server is running on")
		port      = flag.String("PORT", ":8080", "The server port")
		templates = flag.String("TEMPLATES", "templates", "The location of our templates")
	)

	flag.Parse()

	// Generate a basic config to specify the host, port, and template location
	config := ControllerConfig{
		Host:      *host,
		Port:      *port,
		Templates: *templates,
	}

	url := fmt.Sprintf("http://%s%s", config.Host, config.Port)
	fmt.Printf(url)

	// Output: http://localhost:8080
}

func ExampleNewController() {
	// Generate a basic config to specify the host, port, and template location
	config := ControllerConfig{
		Host:      "localhost",
		Port:      ":8080",
		Templates: "templates/",
	}

	// Create a new controller object with the config
	controller := NewController(&config)

	// Create a new Mux handling all routes and attach this to our controller
	mux := http.NewServeMux()
	controller.Attach(mux)
}
