package adminhandlers

import (
	"encoding/json"
	
	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/internal/database/models"
	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
	"github.com/go-playground/validator/v10"
)


func CreateMapHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	
	var reqBody models.CreateMap

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {	
		http.Error(w,"Error decoding request body",http.StatusBadRequest)
		return
	}

	defer	r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(reqBody); err != nil {
		http.Error(w,"validating error",http.StatusBadRequest)
		return
	}
	client := db.NewClient()

	if err := client.Connect(); err != nil {	
		http.Error(w,"Error connecting to database",http.StatusInternalServerError)
		return
	}

	defer client.Disconnect()


	createdMap, err := client.Map.CreateOne(
		db.Map.Name.Set(reqBody.Name),		
		db.Map.Width.Set(reqBody.Width),
		db.Map.Height.Set(reqBody.Height),
		db.Map.Thumbnail.Set(reqBody.Thumbnail),
	).Exec(r.Context())
    
	if err != nil {
		http.Error(w, "error creating map", http.StatusInternalServerError)
		return
	}

	// default map elements creation remains to be done
	

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "map created successfully",
		"mapId" :   createdMap.ID,
	})

	
}