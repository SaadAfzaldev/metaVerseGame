package spacehandlers

import (
	"encoding/json"
	"net/http"


	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
)

// Get space by id
func GetSpacebyId(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")	


	spaceId := r.URL.Query().Get("spaceid")
	
	client := db.NewClient()

	if err := client.Connect(); err != nil {	
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer client.Disconnect()

	space,err := client.Space.FindUnique(db.Space.ID.Equals(spaceId)).
	Exec(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if space == nil {
		http.Error(w, "Space not found", http.StatusNotFound)
		return
		
	}

	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Space found",
		"space": space,
	})

}