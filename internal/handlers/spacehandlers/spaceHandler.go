package spacehandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/internal/database/models"
	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
	"github.com/go-playground/validator/v10"
	
)

func SpaceHandler (w http.ResponseWriter, r * http.Request) {

	w.Header().Set("Content-type","application/json")

	if r.Method != http.MethodPost {
		http.Error(w,"Method Not allowed", http.StatusMethodNotAllowed)
		return
	}

	userId,ok := r.Context().Value("userId").(string)

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	
	var reqBody models.CreateSpace

	
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Fatal("json decoding error")
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
	
	if reqBody.MapId == "" {
		space, err := client.Space.CreateOne(
			db.Space.Name.Set(reqBody.Name),
			db.Space.Width.Set(reqBody.Width),
			db.Space.Height.Set(reqBody.Height),
			db.Space.CreatorID.Set(userId),
		).Exec(r.Context())
		
		if err != nil {
			http.Error(w,"error creating space",http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message" : "space has been created",
			"SpaceId" : space.ID,

		})
	}
			
	
	existingMap, err := client.Map.FindUnique(
		db.Map.ID.Equals(reqBody.MapId),
	).Exec(r.Context())

	if err != nil {
		log.Println("Map not found:", err)
		http.Error(w, "Map not found", http.StatusBadRequest)
		return
	}

	

	space,err := client.Space.CreateOne(
		db.Space.Name.Set(reqBody.Name),
		db.Space.Width.Set(existingMap.Width),	
		db.Space.Height.Set(existingMap.Height),
		db.Space.CreatorID.Set(userId),	
	).Exec(r.Context())

	if err != nil {
		log.Println("Error creating space:", err)
		http.Error(w, "Error creating space", http.StatusInternalServerError)
		return
	}
	
	fmt.Println("space created",space.ID)
	
	  
}