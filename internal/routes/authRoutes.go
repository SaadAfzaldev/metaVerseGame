package routes

import (
	handlers "github.com/SaadAfzaldev/metaVerseGame/internal/handlers/authhandlers"
	"github.com/gorilla/mux"
)


func AuthRoutes(router * mux.Router) {

	apiv1 := router.PathPrefix("/api/v1").Subrouter()

	apiv1.HandleFunc("/signup",handlers.SignupHandler).Methods("POST")
	apiv1.HandleFunc("/signin",handlers.SignInHandler).Methods("POST")
}