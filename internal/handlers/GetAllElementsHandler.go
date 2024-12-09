package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
)

func GetAllElementsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	client := db.NewClient()

	if err := client.Connect(); err != nil {
		http.Error(w,"database connection error", http.StatusInternalServerError)
		return
	}

	defer client.Disconnect()

	elements,err := client.Element.FindMany().Exec(r.Context())	

	if err != nil {
		http.Error(w,"error getting elements", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	
	json.NewEncoder(w).Encode(elements)
}