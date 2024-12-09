package routes

import (
	"github.com/SaadAfzaldev/metaVerseGame/internal/handlers"
	
	"github.com/gorilla/mux"
)

func GetAllElements(router *mux.Router) {
	apiv1 := router.PathPrefix("/api/v1").Subrouter()
	
	apiv1.HandleFunc("/elements",handlers.GetAllElementsHandler).Methods("GET")
}