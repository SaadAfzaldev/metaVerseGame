package spacehandlers

import (
	"encoding/json"
	"net/http"

	
	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
)



func DeleteSpaceHandler (w http.ResponseWriter, r * http.Request) {		

	w.Header().Set("Content-Type","application/json")

	spaceId := r.URL.Query().Get("spaceid")

	userId,ok := r.Context().Value("userId").(string)

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	client := db.NewClient()

	if err := client.Connect(); err != nil {
		http.Error(w,"database connection error",http.StatusInternalServerError)
		return
	}

	defer client.Disconnect()

	space,err:= client.Space.FindUnique(db.Space.ID.Equals(spaceId)).Exec(r.Context())

	
	if err != nil {
		http.Error(w,"space not found",http.StatusNotFound)
		return
	}

	if space.CreatorID != userId {
		http.Error(w,"you are not the owner of this space",http.StatusForbidden)
		return
	}

	// i don't know but you cant delete with findUnique you have to do both separately
	_, err = client.Space.FindUnique(db.Space.ID.Equals(spaceId)).
	Delete().
	Exec(r.Context()) // Chain .Delete() after finding the unique space
	if err != nil {
		http.Error(w, "error deleting space", http.StatusInternalServerError)
		return
	}


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message":"space deleted successfully"})
		
	


}	