package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/register", h.Register).Methods("POST")
	router.HandleFunc("/login", h.Login).Methods("POST")
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if the user already exists
	var existingUser User
	if err := h.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Store the user
	user.Password = string(hashedPassword)
	if err := h.db.Create(&user).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database
	var storedUser User
	if err := h.db.Where("email = ?", user.Email).First(&storedUser).Error; err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
