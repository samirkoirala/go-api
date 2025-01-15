package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Sulav-Adhikari/gouser/pkg/auth"
	"github.com/Sulav-Adhikari/gouser/pkg/models"
	"github.com/Sulav-Adhikari/gouser/pkg/utils"
)

var NewUser models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	// Check if username already exists
	existingUser, err := models.GetUserByUsername(CreateUser.Username)
	if err != nil {
		http.Error(w, "Error checking username", http.StatusInternalServerError)
		return
	}
	if existingUser != nil {
		// Username already exists
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	b := CreateUser.CreateUser()
	res, err := json.Marshal(b)
	if err != nil {
		http.Error(w, "Error while marshaling data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllUsers()
	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "Error while marshaling data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"uname"`
		Password string `json:"password"`
	}

	if err := utils.ParseBody(r, &credentials); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Get the user by username
	user, err := models.GetUserByUsername(credentials.Username)
	if err != nil {
		http.Error(w, "Error checking username", http.StatusInternalServerError)
		return
	}

	// Check if the user exists and if the passwords match
	if user == nil || user.Password != credentials.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// If authentication is successful, generate a JWT
	token, err := auth.GenerateJWT(user.ID, user.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Respond with the JWT token
	response := map[string]string{
		"message": "Login successful",
		"token":   token,
	}

	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUsername(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(map[string]interface{})

	if !ok {
		http.Error(w, "Unauthorized: Unable to extract claims", http.StatusUnauthorized)
		return
	}

	userID := claims["user_id"]
	username := claims["username"]

	response := map[string]interface{}{
		"message":  "JWT is valid. You are authorized to update your username.",
		"user_id":  userID,
		"username": username,
	}
	json.NewEncoder(w).Encode(response)

	// // This will demonstrate that the JWT middleware successfully validated the token
	// w.Header().Set("Content-Type","application/json")

	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("JWT is valid. You are authorized to update username."))
}
