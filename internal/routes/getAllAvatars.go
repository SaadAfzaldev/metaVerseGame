package routes

import (
	"github.com/SaadAfzaldev/metaVerseGame/internal/handlers"
	"github.com/gorilla/mux"
)


func GetAllAvatars(router *mux.Router) {
	apiv1 := router.PathPrefix("/api/v1").Subrouter()
	apiv1.HandleFunc("/avatars",handlers.GetAllAvatarsHandler).Methods("GET")
}