package adminhandlers

import (
	"encoding/json"
	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/internal/database/models"
	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
)

func UpdateElementHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	elementId := r.URL.Query().Get("elementid")

	if elementId == "" {
		http.Error(w, "elementId not found", http.StatusBadRequest)
		return
	}

	var reqBody models.UpdateElement	

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Body Decoding error", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	

	client := db.NewClient()

	if err := client.Connect(); err != nil {
		http.Error(w, "database connection error", http.StatusInternalServerError)
		return
	}

	defer client.Disconnect()

	_, err := client.Element.FindUnique(
		db.Element.ID.Equals(elementId),
	).Update(
		db.Element.ImageURL.Set(reqBody.ImageUrl),
	).Exec(r.Context())

	if err != nil{
		http.Error(w,"error updating element",http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "element updated successfully",
	})
}
