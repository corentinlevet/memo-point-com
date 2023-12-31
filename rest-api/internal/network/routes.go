package network

import (
	"memo-point-com/internal/controllers"
	"memo-point-com/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initAuth(router *gin.Engine) {
	tokenService := &services.TokenService{}

	userService := &services.UserService{
		TokenService: tokenService,
	}

	authService := &services.AuthService{
		UserService: userService,
	}

	authController := &controllers.AuthController{
		AuthService: authService,
		UserService: userService,
	}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/signup", authController.Signup)
	}
}

func initMemo(router *gin.Engine) {
	tokenService := &services.TokenService{}

	userService := &services.UserService{
		TokenService: tokenService,
	}

	memoService := &services.MemoService{}

	memoController := &controllers.MemoController{
		MemoService: memoService,
		UserService: userService,
	}

	memoGroup := router.Group("/memos")
	{
		memoGroup.POST("", memoController.CreateMemo)
		memoGroup.GET("", memoController.GetMemos)
		memoGroup.PUT("", memoController.UpdateMemo)
		memoGroup.DELETE("", memoController.DeleteMemo)
	}
}

func InitRoutes() {
	router := gin.Default()

	cors := cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8081"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Auth-Token"},
	})

	router.Use(cors)

	initAuth(router)
	initMemo(router)

	router.Run(":8080")
}
