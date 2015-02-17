package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jamie-stackhouse/letsgo/Server/models"
)

type ControllerConfig struct {
	Host      string
	Port      string
	Templates string
}

type Controller struct {
	config  *ControllerConfig
	handler *http.Handler
}

// Attach accepts a http.ServeMux object and then creates all routes
func (this *Controller) Attach(mux *http.ServeMux) {
	mux.HandleFunc("/API", this.APIRoute)
	mux.HandleFunc("/", this.DefaultRoute)
}

// Return a new Controller object with its own http.Handler
func NewController(config *ControllerConfig) *Controller {
	return &Controller{
		config:  config,
		handler: new(http.Handler),
	}
}

func (this *Controller) APIRoute(w http.ResponseWriter, r *http.Request) {
	user := models.NewUser("Jeff")
	userData, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(userData)
}

func (this *Controller) DefaultRoute(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://%s%s/API", this.config.Host, this.config.Port)
	userData, err := http.Get(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(userData.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.NewUser("Jeff")
	json.Unmarshal(body, &user)

	filename := "index.html"
	templatePath := fmt.Sprintf("%s/%s", this.config.Templates, filename)
	template, err := template.ParseFiles(templatePath)

	if err != nil {
		log.Fatal(err)
	}

	template.Execute(w, user)
}
