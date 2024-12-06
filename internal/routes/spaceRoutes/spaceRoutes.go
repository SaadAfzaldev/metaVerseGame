package spaceRoutes

import (
	"github.com/SaadAfzaldev/metaVerseGame/internal/handlers/spacehandlers"
	"github.com/SaadAfzaldev/metaVerseGame/internal/middlewares"
	"github.com/gorilla/mux"
)

func SetUpSpaceRoutes (router * mux.Router) {

	apiv1 := router.PathPrefix("/api/v1").Subrouter()
	spaceRouter:= apiv1.PathPrefix("/space").Subrouter()
	spaceRouter.Use(middlewares.UserMiddleware)

	spaceRouter.HandleFunc("/",spacehandlers.SpaceHandler).Methods("POST")
	spaceRouter.HandleFunc("/:spaceid",spacehandlers.DeleteSpaceHandler).Methods("DELETE")
	spaceRouter.HandleFunc("/all",spacehandlers.GetAllSpacesHandler).Methods("GET")
	spaceRouter.HandleFunc("/element",spacehandlers.AddElementHandler).Methods("POST")
	spaceRouter.HandleFunc("/element",spacehandlers.AddElementHandler).Methods("DELETE")
}
