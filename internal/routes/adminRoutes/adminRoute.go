package adminroutes

import (
	"github.com/SaadAfzaldev/metaVerseGame/internal/handlers/adminhandlers"
	"github.com/SaadAfzaldev/metaVerseGame/internal/middlewares"
	"github.com/gorilla/mux"
)



func SetUpAdminRoutes (router * mux.Router) {
	apiv1 := router.PathPrefix("/api/v1").Subrouter()
	adminRouter := apiv1.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middlewares.AdminMiddleware)
	adminRouter.HandleFunc("/element",adminhandlers.AddElementHandler).Methods("POST")
	adminRouter.HandleFunc("/element/",adminhandlers.UpdateElementHandler).Methods("PUT")
	adminRouter.HandleFunc("/avatar",adminhandlers.CreateAvatarHandler).Methods("POST")
	adminRouter.HandleFunc("/createmap",adminhandlers.CreateMapHandler).Methods("POST")
}