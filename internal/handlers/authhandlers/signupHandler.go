package handlers

import (
	"encoding/json"

	"log"

	"net/http"

	
	"github.com/SaadAfzaldev/metaVerseGame/internal/database/models"
	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)
func SignupHandler (w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost {
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		return
	}

	var signUpBody models.UserSignUpBody

	if err := json.NewDecoder(r.Body).Decode(&signUpBody); err != nil {
		log.Fatal("json decoding error")
		http.Error(w,"json decoding error",http.StatusBadRequest)
		return
	
	} 

	defer r.Body.Close()

	validate := validator.New()

	if err := validate.Struct(signUpBody); err != nil {			
		log.Fatal("invalid input")
		http.Error(w,"invalid input",http.StatusBadRequest)
		return
	}
	
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		log.Printf("Error connecting to the database: %v", err)
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	} 
	defer client.Disconnect() 

	existingUser, err := client.User.FindUnique(
	db.User.Username.Equals(signUpBody.Username),
	).Exec(r.Context())

	if err == nil && existingUser != nil {

		http.Error(w,"User already exist", http.StatusBadRequest)
		return
	}

	role := db.RoleUser
	if signUpBody.Role != "" {

		switch signUpBody.Role {

		case "Admin" :
			role = db.RoleAdmin
		case "User" :
			role = db.RoleUser
		default:
			http.Error(w, "Invalid Role", http.StatusBadRequest)
			return	
		}

	
	}
	
	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(signUpBody.Password),bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("Error hashing password", err)
		return
	}
	password := string(hashedPassword)
	
	user, err := client.User.CreateOne(
		db.User.Username.Set(signUpBody.Username),
		db.User.Password.Set(password),
		db.User.Role.Set(role),
	).Exec(r.Context())

	if err != nil {

		log.Fatal("Error creating user",err)
		http.Error(w,"Error creating user",http.StatusBadRequest)
		return

	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "user created successfully",
		"userId is" : user.ID,
	})
}
