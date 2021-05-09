package controller

import (
	"fmt"
	"net/http"

	"github.com/DarioAle/counter-api"
)

// Controller handles the HTTP requests
type Controller interface {
	Healthz(w http.ResponseWriter, r *http.Request)
	Statusz(w http.ResponseWriter, r *http.Request)

	CreateLog(w http.ResponseWriter, r *http.Request)
	UpdateLog(w http.ResponseWriter, r *http.Request)
}

// controller struct holds the usecase
type controller struct {
	commit  string
	version string

	repository counter.Repository
	storage    counter.Storage
}

// New returns a controller
func New(
	c string,
	v string,
	r counter.Repository,
	s counter.Storage,
) Controller {
	fmt.Println("Init controller")
	return &controller{
		commit:     c,
		version:    v,
		repository: r,
		storage:    s,
	}
}
