package spacehandlers

import (
	"encoding/json"
	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/internal/database/models"
	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
	"github.com/go-playground/validator/v10"
)


func AddElementHandler (w http.ResponseWriter, r * http.Request) {

	w.Header().Set("Content-Type","application/json")


	var reqBody models.AddElement

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w,"json decoding error",http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(reqBody); err != nil {
		http.Error(w,"validating error", http.StatusBadRequest)
		return
	}

	client := db.NewClient()

	if err := client.Connect();err != nil {
		http.Error(w,"database connection error", http.StatusBadRequest)
		return
	}

	defer client.Disconnect()

	space, err := client.Space.FindUnique(
		db.Space.ID.Equals(reqBody.SpaceId),
	).Select(
		db.Space.Width.Field(), 
		db.Space.Height.Field(),
	).Exec(r.Context())

	if err != nil {	
		http.Error(w,"space not found",http.StatusNotFound)
		return
	}

	if  space.CreatorID!= r.Context().Value("userId").(string) {
		http.Error(w,"you are not the owner of this space",http.StatusForbidden)
		return
	}
	

	width, _ := space.Width()
	height, _ := space.Height()

	if reqBody.X > width || reqBody.Y > height || reqBody.X < 0 || reqBody.Y < 0 {
		http.Error(w,"element out of bounds",http.StatusBadRequest)
		return
		
	}

	
	client.SpaceElements.CreateOne(
		db. SpaceElements.SpaceID.Set(reqBody.SpaceId),
		db.SpaceElements.ElementID.Set(reqBody.ElementId),
		db.SpaceElements.X.Set(reqBody.X),
		db.SpaceElements.Y.Set(reqBody.Y),
	).Exec(r.Context())

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message":"element added successfully"})
}

