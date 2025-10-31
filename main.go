package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	r := gin.Default()

	// front := os.Getenv("FRONTEND_URL")
	// if front == "" {
	// 	front = "http://localhost:3000"
	// }
	front := "https://front-end-next-js-v1-gj6b2uz5m-satriaan-jpas-projects.vercel.app/"

	origins := strings.Split(front, ",")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "PING!!!"})
	})

	r.GET("/apps", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome To Task Manager"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
