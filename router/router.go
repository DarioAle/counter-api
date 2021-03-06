package router

import (
	"github.com/DarioAle/counter-api/controller"

	"github.com/gorilla/mux"
)

// New returns a *mux.Router
func New(
	app controller.Controller,
) *mux.Router {
	router := mux.NewRouter()

	setupLogController(router, app)
	setupAppController(router, app)

	return router
}
