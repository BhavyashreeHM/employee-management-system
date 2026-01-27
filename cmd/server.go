package main

import (
	"employee-management-system/internal/api/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"
	r := gin.Default()
	r.GET("/", handlers.Emploee)
	log.Println("Server running on port", port)
	r.Run(":" + port)
}
