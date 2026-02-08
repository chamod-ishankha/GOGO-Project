package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/middleware"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/model"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/repository"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/pkg/utils"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register endpoint hit")

	var req model.Register
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	// ✅ Check if email already exists
	existingUser, _ := h.Repo.GetUserByEmail(req.Email)
	if existingUser != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Email already exists"})
		return
	}

	// ✅ Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not hash password"})
		return
	}
	req.Password = hashedPassword

	// ✅ Save user
	if err := h.Repo.CreateUser(&req); err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	// ✅ Generate JWT token
	token, _ := utils.GenerateJWT(req.ID, req.Role)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login endpoint hit")
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Invalid login request: %v", err)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	user, err := h.Repo.GetUserByEmail(req.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("Login failed for email %s: %v", req.Email, err)
		json.NewEncoder(w).Encode(map[string]string{"error": "Email or password incorrect"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("Password mismatch for email %s", req.Email)
		json.NewEncoder(w).Encode(map[string]string{"error": "Email or password incorrect"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Token generation failed for user ID %d: %v", user.ID, err)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not generate token"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(*utils.Claims)

	user, err := h.Repo.GetUserByID(claims.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateMe(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(*utils.Claims)

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Name and email required"})
		return
	}

	err := h.Repo.UpdateUser(claims.UserID, req.Name, req.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Profile updated successfully",
	})
}
