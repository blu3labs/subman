package server

import (
	"log"
	"net/http"
	router "subman/api/router"

	"github.com/rs/cors"
)

func Start() {
	r := router.Router()
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "PUT", "POST", "PATCH"},
		AllowedHeaders: []string{"a_custom_header", "content_type"},
	}).Handler(r)
	log.Fatal(http.ListenAndServe(":10203", handler))
}