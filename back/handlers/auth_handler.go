package handlers

import (
	"context"
	"log"
	"net/http"
	"ped_poject/database"
	"ped_poject/models"
	"ped_poject/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserRepo *database.UserRepository
}

func NewAuthHandler(repo *database.UserRepository) *AuthHandler {
	return &AuthHandler{UserRepo: repo}
}

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"" binding:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("invalid request: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := h.UserRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Printf("database error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	if user == nil {
		log.Printf("user %v not found\n", req.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		log.Printf("wrong password for user %v\n", req.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		log.Printf("generation token error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "generation token error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"auth_token": token})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("invalid request: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	existingUser, _ := h.UserRepo.GetUserByEmail(context.Background(), req.Email)
	if existingUser != nil {
		log.Printf("user with this username exists\n")
		c.JSON(http.StatusConflict, gin.H{"error": "user with this username already exists"})
		return
	}
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Printf("hashing password error: %v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
		return
	}
	user := models.User{
		Email:        req.Email,
		PasswordHash: hash,
		Name:         req.Name,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = h.UserRepo.CreateUser(ctx, &user)
	if err != nil {
		log.Printf("creating user error: %v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "creating user error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
