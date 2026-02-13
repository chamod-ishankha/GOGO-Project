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
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	existingUser, _ := h.Repo.GetUserByEmail(req.Email)
	if existingUser != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Email already exists")
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to secure password")
		return
	}
	req.Password = hashedPassword

	if err := h.Repo.CreateUser(&req); err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	token, err := utils.GenerateJWT(req.ID, req.Role)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := map[string]interface{}{
		"message": "User registered successfully",
		"token":   token,
	}

	utils.WriteJSONResponse(w, http.StatusCreated, response)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login endpoint hit")

	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Invalid login request: %v", err)
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := h.Repo.GetUserByEmail(req.Email)
	if err != nil {
		log.Printf("Login failed for %s: %v", req.Email, err)
		utils.WriteJSONError(w, http.StatusUnauthorized, "Email or password incorrect")
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		log.Printf("Password mismatch for %s", req.Email)
		utils.WriteJSONError(w, http.StatusUnauthorized, "Email or password incorrect")
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		log.Printf("Token generation failed for user %d: %v", user.ID, err)
		utils.WriteJSONError(w, http.StatusInternalServerError, "Could not generate token")
		return
	}

	response := map[string]interface{}{
		"token": token,
		"user": map[string]string{
			"username": user.Email,
			"role":     user.Role,
		},
	}

	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetMe endpoint hit")

	claims, ok := r.Context().Value(middleware.UserContextKey).(*utils.Claims)
	if !ok {
		utils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.Repo.GetUserByID(claims.UserID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	response := user

	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *UserHandler) UpdateMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateMe endpoint hit")

	claims, ok := r.Context().Value(middleware.UserContextKey).(*utils.Claims)
	if !ok {
		utils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Name == "" || req.Email == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, "Name and email are required")
		return
	}

	if err := h.Repo.UpdateUser(claims.UserID, req.Name, req.Email); err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	response := map[string]string{
		"message": "Profile updated successfully",
	}

	utils.WriteJSONResponse(w, http.StatusOK, response)
}
