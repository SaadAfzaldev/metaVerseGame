package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/internal/routes"
	adminroutes "github.com/SaadAfzaldev/metaVerseGame/internal/routes/adminRoutes"
	"github.com/SaadAfzaldev/metaVerseGame/internal/routes/spaceRoutes"
	"github.com/SaadAfzaldev/metaVerseGame/internal/routes/userRoutes"
	"github.com/gorilla/mux"
)

func main() {

	
	

	router := mux.NewRouter()

	routes.AuthRoutes(router)
	routes.GetAllElements(router)

	userroutes.SetupUserRoutes(router)
	spaceRoutes.SetUpSpaceRoutes(router)
	adminroutes.SetUpAdminRoutes(router)

	
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
	
}
