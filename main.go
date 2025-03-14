package main

import (
	"chat-app/controllers"
	"chat-app/initializers"
	"chat-app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.Static("/public", "./public")
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"Title": "Home"})
	})
	r.GET("/chat", middleware.RequireAuth, func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", gin.H{"Title": "Chat Room"})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{"Title": "Login"})
	})
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{"Title": "Signup"})
	})

	api := r.Group("/api")
	{
		api.POST("/signup", controllers.Signup)
		api.POST("/login", controllers.Login)
		api.GET("/validate", middleware.RequireAuth, controllers.Validate)
		//TODO : api.GET("/chat/ws", controllers.ChatWebSocket)
	}

	r.Run(":8080")
}
