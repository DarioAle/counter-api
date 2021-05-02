package router

import (
	"github.com/DarioAle/counter-api/controller"

	"github.com/gorilla/mux"
)

// setupLogController setup the router with the log controller
func setupLogController(router *mux.Router, cont controller.Controller) {
	router.HandleFunc("/logs", cont.CreateLog).Methods("POST")
	router.HandleFunc("/logs/{id}", cont.UpdateLog).Methods("PATCH")
}
