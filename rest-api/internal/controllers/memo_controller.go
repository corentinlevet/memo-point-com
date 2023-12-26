package controllers

import (
	"memo-point-com/internal/services"

	"github.com/gin-gonic/gin"
)

type MemoController struct {
	MemoService *services.MemoService
	UserService *services.UserService
}

// POST /memos
func (controller *MemoController) CreateMemo(ctx *gin.Context) {
	if ctx.Request.Method != "POST" {
		ctx.JSON(405, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	authToken := ctx.GetHeader("Auth-Token")
	if authToken == "" {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := controller.UserService.GetUserByAuthToken(authToken)
	if err != nil || user == nil {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var json struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	err = controller.MemoService.CreateMemo(user.ID, json.Title, json.Content)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Memo created",
	})
}

// GET /memos
func (controller *MemoController) GetMemos(ctx *gin.Context) {
	if ctx.Request.Method != "GET" {
		ctx.JSON(405, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	authToken := ctx.GetHeader("Auth-Token")
	if authToken == "" {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := controller.UserService.GetUserByAuthToken(authToken)
	if err != nil || user == nil {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	memos, err := controller.MemoService.GetMemosByUserID(user.ID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"memos": memos,
	})
}

// PUT /memos
func (controller *MemoController) UpdateMemo(ctx *gin.Context) {
	if ctx.Request.Method != "PUT" {
		ctx.JSON(405, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	authToken := ctx.GetHeader("Auth-Token")
	if authToken == "" {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := controller.UserService.GetUserByAuthToken(authToken)
	if err != nil || user == nil {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var json struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	err = controller.MemoService.UpdateMemoByID(user.ID, json.ID, json.Title, json.Content)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Memo updated",
	})
}

// DELETE /memos
func (controller *MemoController) DeleteMemo(ctx *gin.Context) {
	if ctx.Request.Method != "DELETE" {
		ctx.JSON(405, gin.H{
			"message": "Method not allowed",
		})
		return
	}

	authToken := ctx.GetHeader("Auth-Token")
	if authToken == "" {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, err := controller.UserService.GetUserByAuthToken(authToken)
	if err != nil || user == nil {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var json struct {
		ID int `json:"id"`
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	err = controller.MemoService.DeleteMemoByID(user.ID, json.ID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Memo deleted",
	})
}
