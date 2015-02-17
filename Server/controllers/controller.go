package controller

import (
	"net/http"
)

type Controller struct{}

func (this *Controller) Attach(mux *http.ServeMux) {
}

func NewController() *Controller {
	return &Controller{}
}
