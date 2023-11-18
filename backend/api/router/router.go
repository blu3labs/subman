package router

import (
	"subman/api/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/executeables", controller.GetExecuteableSubscriptions).Methods("GET")

	router.HandleFunc("/subscription", controller.PutSubscription).Methods("POST")

	return router
}
