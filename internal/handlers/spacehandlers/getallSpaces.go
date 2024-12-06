package spacehandlers

import (
	"encoding/json"
	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
)


func GetAllSpacesHandler (w http.ResponseWriter, r * http.Request) {

	w.Header().Set("Content-Type","application/json")

	client := db.NewClient()
	if err := client.Connect(); err != nil {
		http.Error(w,"database connection error",http.StatusInternalServerError)
		return
	}

	defer client.Disconnect()

	spaces,err := client.Space.FindMany(
		db.Space.CreatorID.Equals(r.Context().Value("userId").(string)),
	).Exec(r.Context())

	if err != nil {
		http.Error(w,"error getting spaces",http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(spaces)

}