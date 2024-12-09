package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
)

func GetAllAvatarsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	client := db.NewClient()

	if err := client.Connect(); err != nil {
		http.Error(w, "database connection error", http.StatusInternalServerError)
		return
	}

	defer client.Disconnect()

	avatars, err := client.Avatar.FindMany().Exec(r.Context())

	if err != nil {
	    http.Error(w, "error retrieving avatars", http.StatusInternalServerError)
	    return
	}

	json.NewEncoder(w).Encode(avatars)
}