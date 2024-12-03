package userHandler

import (
	"encoding/json"

	"log"

	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/internal/database/models"
	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
	"github.com/go-playground/validator/v10"
	
)

func MetaDataHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type","application/json")

	userId,ok := r.Context().Value("userId").(string)	

	if !ok {
		http.Error(w,"Unauthorized",http.StatusUnauthorized)
		return
	}

	

	var metaDataUpdateBody models.MetaDataUpdate

	if err := json.NewDecoder(r.Body).Decode(&metaDataUpdateBody); err != nil {
		log.Fatal("json decoding error")
		http.Error(w,"json decoding error",http.StatusBadRequest)
		return
	}

	validate := validator.New()

	if err := validate.Struct(metaDataUpdateBody); err != nil {
		http.Error(w,"Error validating body",http.StatusBadRequest)
		return
	}  
	

	client := db.NewClient()
	
	if err := client.Connect(); err != nil {
		http.Error(w,"Error connecting database",http.StatusBadRequest)
		return
	}

	defer client.Disconnect()

	_, err := client.User.FindUnique(
		db.User.ID.Equals(userId),
	).Update(
		db.User.AvatarID.Set(metaDataUpdateBody.AvatarId),
	).Exec(r.Context())

	if err != nil {
		http.Error(w, "Failed to update metadata", http.StatusInternalServerError)
            return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message" : "User's metadata updated successfully",
	})

}
 