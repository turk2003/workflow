package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "github.com/turk2003/workflow/config"
	"github.com/turk2003/workflow/internal/controllers"
	repositories "github.com/turk2003/workflow/internal/respositories"
	"github.com/turk2003/workflow/internal/services"
)

func main() {
	db, err := gorm.Open(
		postgres.Open(
			"postgres://postgres:password@localhost:5432/task5",
		),
	)
	if err != nil {
		log.Panic(err)
	}
	// db := config.SetupDatabase()
	itemRepo := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService)

	router := gin.Default()

	router.POST("/items", itemController.CreateItem)
	router.GET("/items", itemController.GetAllItems)       // Get all items
	router.GET("/items/:id", itemController.GetItemByID)   // Get an item by ID
	router.PUT("/items/:id", itemController.UpdateItem)    // Update an item by ID
	router.DELETE("/items/:id", itemController.DeleteItem) // Delete an item by ID

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	router.Run(port)
}
