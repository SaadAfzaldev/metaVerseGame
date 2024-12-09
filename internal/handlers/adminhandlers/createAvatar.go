package adminhandlers

import (
	"encoding/json"
	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/internal/database/models"
	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
	"github.com/go-playground/validator/v10"
	
)


func CreateAvatarHandler (w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var reqBody models.CreateAvatar
	
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {	
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	validate := validator.New()


	if err := validate.Struct(reqBody); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}


	client := db.NewClient()
	
	if err := client.Connect(); err != nil {
		http.Error(w, "error connecting to database", http.StatusInternalServerError)
		return
	}

	defer client.Disconnect()

	avatar,err := client.Avatar.CreateOne(
		db.Avatar.ImageURL.Set(reqBody.ImageUrl),
		db.Avatar.Name.Set(reqBody.Name),
	).Exec(r.Context())

	if err != nil {
		http.Error(w, "error creating avatar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "avatar created successfully",
		"avatar": avatar.ID,
	})
}