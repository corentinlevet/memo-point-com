package controllers

import (
	"memo-point-com/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
	UserService *services.UserService
}

// POST /auth/login
func (controller *AuthController) Login(ctx *gin.Context) {
	if ctx.Request.Method != "POST" {
		ctx.JSON(405, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	var json struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	user, err := controller.AuthService.Login(json.Username, json.Password)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Login successful",
		"user_id": user.ID,
	})
}

// POST /auth/signup
func (controller *AuthController) Signup(ctx *gin.Context) {
	if ctx.Request.Method != "POST" {
		ctx.JSON(405, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	var json struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	user, err := controller.AuthService.Signup(json.Username, json.Email, json.Password)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Signup successful",
		"user_id": user.ID,
	})
}
