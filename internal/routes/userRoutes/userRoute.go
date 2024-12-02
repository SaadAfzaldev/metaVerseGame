package userroutes

import (
	"github.com/SaadAfzaldev/metaVerseGame/internal/handlers/userhandlers"
	"github.com/SaadAfzaldev/metaVerseGame/internal/middlewares"
	"github.com/gorilla/mux"
)




func  SetupUserRoutes (router * mux.Router)  {

	apiv1 := router.PathPrefix("/api/v1").Subrouter()
	userRouter := apiv1.PathPrefix("/user").Subrouter()

	userRouter.Use(middlewares.AuthMiddleware)

	userRouter.HandleFunc("/metadata",userHandler.MetaDataHandler).Methods("POST")
}

