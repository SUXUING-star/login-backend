package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"login-backend/config"
	"login-backend/db"
	"login-backend/handlers"
	"login-backend/middleware"
)

func init() {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// 加载配置
	if err := config.Init(); err != nil {
		log.Fatal("Failed to load config:", err)
	}
}

func main() {
	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 初始化数据库连接
	if err := db.Init(ctx); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer db.Close(ctx)

	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	// 初始化路由
	router := setupRouter()

	// 启动服务器
	port := config.GlobalConfig.Server.Port
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// 中间件
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())

	// API路由
	api := r.Group("/api")
	{
		// 认证相关
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
			auth.GET("/verify-email", handlers.VerifyEmail)
			auth.POST("/forgot-password", handlers.ForgotPassword)
			auth.POST("/reset-password", handlers.ResetPassword)
		}

		// 需要认证的路由
		authorized := api.Group("")
		authorized.Use(middleware.Auth())
		{
			// 用户相关
			user := authorized.Group("/user")
			{
				user.PUT("/password", handlers.UpdatePassword)
			}
		}
	}
	return r
}
