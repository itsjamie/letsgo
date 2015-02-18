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

// ControllerConfig allows for storage of the host, port, and template location
type ControllerConfig struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	Templates string `json:"-"`
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

// APIRoute creates a new User object with the username 'jeff' then returns the marshalled object
func (this *Controller) APIRoute(w http.ResponseWriter, r *http.Request) {
	// create a new User object for 'jeff'
	user := models.NewUser("Jeff")

	// marshal the object to JSON
	userData, err := json.Marshal(user)

	// incase of error return a internal server error (500) and exit the function
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write the JSON object to the response
	w.Write(userData)
}

func (this *Controller) DefaultRoute(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://%s%s/API", this.config.Host, this.config.Port)
	userData, err := http.Get(url)

	// incase of error return a internal server error (500) and exit the function
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(userData.Body)

	// incase of error return a internal server error (500) and exit the function
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
