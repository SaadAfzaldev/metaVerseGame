package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	
	"github.com/SaadAfzaldev/metaVerseGame/internal/database/models"
	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignInHandler (w http.ResponseWriter, r * http.Request) {

	w.Header().Set("Content-type", "application/json")

	if r.Method != http.MethodPost {

		http.Error(w,"Invalid Method", http.StatusMethodNotAllowed)
		return 
	}
	var signinBody models.UserSignInBody

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

	client := db.NewClient()

	if err := client.Connect(); err != nil {
		log.Printf("Error connecting to the database: %v", err)
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	} 
		
	defer client.Disconnect()

	user,err := client.User.FindUnique(
		db.User.Username.Equals(signinBody.Username),
	).Exec(r.Context())

	if err != nil {
		log.Printf("Error finding user in database: %v", err)
		http.Error(w, "Error finding user in database or user already exist", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(signinBody.Password)); err != nil {
		log.Printf("Error validating password: %v", err)
		http.Error(w, "Invalid credentials", http.StatusInternalServerError)
		return
	}

	claims := jwt.MapClaims{
		
		"userId" : user.ID,
		"role"  : user.Role,

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	var jwtSecret = os.Getenv("JWT_SECRET") 

	if jwtSecret == "" {
		http.Error(w, "JWT_SECRET not set", http.StatusInternalServerError)
		return
	  }

	signedToken, err := token.SignedString([]byte(jwtSecret)) 

	if err != nil {
		http.Error(w,"error signing jwt",http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"token" : signedToken,
	} 

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	
}