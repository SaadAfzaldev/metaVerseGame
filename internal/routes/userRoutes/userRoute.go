package userroutes

import (
	"github.com/SaadAfzaldev/metaVerseGame/internal/handlers/userhandlers"
	"github.com/gorilla/mux"
)




func  SetupUserRoutes (router * mux.Router)  {

	apiv1 := router.PathPrefix("/api/v1").Subrouter()
	userRouter := apiv1.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/profile",userHandler.ProfileHandler).Methods("GET")
}