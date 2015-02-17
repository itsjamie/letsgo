package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jamie-stackhouse/letsgo/Server/models"
)

type Controller struct {
	handler *http.Handler
}

// Attach accepts a http.ServeMux object and then creates all routes
func (this *Controller) Attach(mux *http.ServeMux) {
	mux.HandleFunc("/", this.DefaultRoute)
}

// Return a new Controller object with its own http.Handler
func NewController() *Controller {
	return &Controller{handler: new(http.Handler)}
}

func (this *Controller) DefaultRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s Requested\n", r.URL)

	user := models.NewUser("jeff")

	template, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatal(err)
	}

	template.Execute(w, user)
}
