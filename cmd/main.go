package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SaadAfzaldev/metaVerseGame/internal/routes"
	"github.com/SaadAfzaldev/metaVerseGame/internal/routes/userRoutes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	databaseURL := os.Getenv("DATABASE_URL");
	

	if databaseURL == "" {
		log.Fatal("database url must be set")
	}
		 
	fmt.Println("DATABASE_URL:", databaseURL)

	router := mux.NewRouter()

	routes.AuthRoutes(router)

	userroutes.SetupUserRoutes(router)

	
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
	
}
