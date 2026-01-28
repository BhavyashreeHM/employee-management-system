package main

import (
	"employee-management-system/internal/api/handlers"
	"employee-management-system/internal/reposiory/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"
	r := gin.Default()

	db, err := database.ConnectSQL()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	r.GET("/", handlers.Employee)
	log.Println("Server running on port", port)
	r.Run(":" + port)
}
