package adminhandlers

import (
	"encoding/json"
	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/internal/database/models"
	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
	"github.com/go-playground/validator/v10"
)


func AddElementHandler (w http.ResponseWriter, r * http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var reqBody models.CreateElement

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {	
		http.Error(w, "Body Decoding error", http.StatusBadRequest)
		return

	}

	defer r.Body.Close()

	validate := validator.New()

	if err := validate.Struct(reqBody); err != nil {
		http.Error(w, "Validation Error", http.StatusBadRequest)
		return
		
	}

	client := db.NewClient()

	if err := client.Connect(); err != nil {
		http.Error(w, "database connection error", http.StatusInternalServerError)
		return
	}
	
	defer client.Disconnect()

	element,err := client.Element.CreateOne(
		db.Element.Width.Set(reqBody.Width),
		db.Element.Height.Set(reqBody.Height),
		db.Element.ImageURL.Set(reqBody.ImageUrl),
		db.Element.Static.Set(reqBody.Status),
	).Exec(r.Context())

	if err != nil {
		http.Error(w,"error creating element", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "element created successfully",
		"elementId": element.ID,
	})
}
