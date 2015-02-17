package controllers

import (
	"io"
	"net/http"
)

type Controller struct{}

func (this *Controller) Attach(mux *http.ServeMux) {
	mux.HandleFunc("/", this.DefaultRoute)
	mux.HandleFunc("/hello", this.Hello)
}

func NewController() *Controller {
	return &Controller{}
}

func (this *Controller) DefaultRoute(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world\n")
}

func (this *Controller) Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world\n")
}
