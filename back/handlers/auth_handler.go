package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
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
	Email          string `json:"email" binding:"required,email"`
	Name           string `json:"name" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Check_Password string `json:"check_password" binding:"required"`
}
type UpdateProfileRequest struct {
	Email          string `json:"email" binding:"required,email"`
	New_Name       string `json:"new_name" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Check_Password string `json:"check_password" binding:"required"`
}

type Check_EmailRequest struct {
	Email string `json:"email" binding:"required,email"`
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
	if req.Check_Password != req.Password {
		log.Printf("passwords must match\n")
		c.JSON(http.StatusBadRequest, gin.H{"error": "passwords must match"})
		return

	}
	existingUser_name, _ := h.UserRepo.GetUserByName(context.Background(), req.Name)
	if existingUser_name != nil {
		log.Printf("user with this name exists\n")
		c.JSON(http.StatusConflict, gin.H{"error": "user with this name already exists"})
		return
	}
	existingUser, _ := h.UserRepo.GetUserByEmail(context.Background(), req.Email)
	if existingUser != nil {
		log.Printf("user with this email exists\n")
		c.JSON(http.StatusConflict, gin.H{"error": "user with this email already exists"})
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
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	var req UpdateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("invalid request: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, _ := h.UserRepo.GetUserByEmail(context.Background(), req.Email)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	existingUser, _ := h.UserRepo.GetUserByName(context.Background(), req.New_Name)
	if existingUser != nil && existingUser.Email != req.Email {
		c.JSON(http.StatusConflict, gin.H{"error": "user with this name already exists"})
		return
	}
	if req.New_Name != "" {
		user.Name = req.New_Name
	}
	if req.Password != req.Check_Password {
		log.Printf("passwords must match\n")
		c.JSON(http.StatusBadRequest, gin.H{"error": "passwords must match"})
		return
	}
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Printf("hashing password error: %v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
		return
	}
	user.PasswordHash = hash
	err = h.UserRepo.UpdateUser(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (h *AuthHandler) CheckEmail(c *gin.Context) {
	var req Check_EmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("invalid request: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	existing_User, err := h.UserRepo.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		log.Printf("database error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	if existing_User == nil {
		log.Printf("email %v not found\n", req.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email not found"})
		return
	}
	user := "5c5516c6b12e23"
	password := "375e7c6f6044b7"
	from := "algorithmtrainerinc@gmail.com"
	to := []string{req.Email}
	addr := "smtp.mailtrap.io:2525"
	host := "smtp.mailtrap.io"
	msg := []byte("From: " + from + "\r\n" +
		"To: " + req.Email + "\r\n" +
		"Subject: Test mail\r\n\r\n" +
		"Email body\r\n")
	auth := smtp.PlainAuth("", user, password, host)
	err = smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
}
