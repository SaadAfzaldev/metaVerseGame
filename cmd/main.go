package main

import (
	"fmt"
	"log"
	"net/http"
	

	"github.com/SaadAfzaldev/metaVerseGame/internal/routes"
	"github.com/SaadAfzaldev/metaVerseGame/internal/routes/spaceRoutes"
	"github.com/SaadAfzaldev/metaVerseGame/internal/routes/userRoutes"
	"github.com/gorilla/mux"
	
)

func main() {

	
	

	router := mux.NewRouter()

	routes.AuthRoutes(router)

	userroutes.SetupUserRoutes(router)
	spaceRoutes.SetUpSpaceRoutes(router)

	
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
	
}
