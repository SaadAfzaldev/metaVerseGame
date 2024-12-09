package spacehandlers

import (
	"encoding/json"
	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/internal/database/models"
	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
	"github.com/go-playground/validator/v10"
	
)


func DeleteElementHandler (w http.ResponseWriter, r * http.Request) {
	

	userId,ok := r.Context().Value("userId").(string)

	if !ok {
		http.Error(w,"Unauthorized",http.StatusUnauthorized)
		return
	}

	var reqBody models.DeleteElement
	
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w,"Failed to decode request body",http.StatusBadRequest)
		return
	}
	
	validate := validator.New()

	if err := validate.Struct(reqBody); err != nil {
		http.Error(w,"Failed to validate request body",http.StatusBadRequest)
		return
	}

	client := db.NewClient()

	if err := client.Connect(); err != nil {
		http.Error(w,"Error connecting database", http.StatusBadRequest)
		return
	}

	defer client.Disconnect()

	space,err := client.SpaceElements.FindUnique(
		db.SpaceElements.ID.Equals(reqBody.ElementId),
	).Exec(r.Context())


	if err != nil {
		http.Error(w,"element not found ",http.StatusBadRequest)
		return
	}

	if space.Space().CreatorID != userId {
		http.Error(w,"You are not the owner of this space",http.StatusBadRequest)
		return
	}

	
	_,err = client.SpaceElements.FindUnique(
		db.SpaceElements.ID.Equals(reqBody.ElementId),
	).Delete().Exec(r.Context())	

	if err != nil {
		http.Error(w,"Error deleting element",http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	
	json.NewEncoder(w).Encode(map[string]string{"message":"Element deleted successfully"})

}
