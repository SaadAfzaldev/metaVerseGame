package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SaadAfzaldev/metaVerseGame/internal/database"
	"github.com/go-playground/validator/v10"
)

func SignInHandler (w http.ResponseWriter, r * http.Request) {

	w.Header().Set("Content-type", "application/json")

	if r.Method != http.MethodPost {

		http.Error(w,"Invalid Method", http.StatusMethodNotAllowed)
		return 
	}
	var signinBody database.UserSignInBody 

	if err := json.NewDecoder(r.Body).Decode(&signinBody); err != nil {
		
		http.Error(w, "Decoding Error", http.StatusBadRequest)
        	return
	} 

	defer r.Body.Close()

	validate :=  validator.New()
	if err := validate.Struct(signinBody); err != nil {
		http.Error(w, "Validation  Error", http.StatusBadRequest)
        	return
	}
	
	response := map[string]interface{}{
		"message": "User sign in successfully",
		
	}
    
	 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}