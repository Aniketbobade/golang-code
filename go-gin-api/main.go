package main


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create a Gin router
	router := gin.Default()

	// Routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin REST API 🚀",
		})
	})

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"users": []string{"Aniket", "Sagar", "Priya"},
		})
	})

	// Start server on port 8080
	router.Run(":8080")
}
