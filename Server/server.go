package main

import (
	"github.com/jamie-stackhouse/letsgo/Server/controller"
	"io"
	"log"
	"net/http"
)

var (
	port       string      = ":8080"
	controller *Controller = NewController()
)

func (this *Controller) DefaultRoute(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world\n")
}

func (this *Controller) Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world\n")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.DefaultRoute)
	mux.HandleFunc("/hello", controller.Hello)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
