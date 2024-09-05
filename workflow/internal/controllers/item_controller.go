package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	constants "github.com/turk2003/workflow/internal/constant"
	"github.com/turk2003/workflow/internal/models"
	"github.com/turk2003/workflow/internal/services"
)

type ItemController struct {
	Service services.ItemService
}

func NewItemController(service services.ItemService) *ItemController {
	return &ItemController{Service: service}
}

func (ctrl *ItemController) CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default status to PENDING if not provided
	if item.Status == "" {
		item.Status = constants.ItemPendingStatus
	}

	if err := ctrl.Service.CreateItem(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

// Add more controller methods (GetAllItems, GetItemByID, UpdateItem, DeleteItem) similarly
// GetAllItems retrieves all items
func (ctrl *ItemController) GetAllItems(c *gin.Context) {
	items, err := ctrl.Service.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// GetItemByID retrieves a single item by ID
func (ctrl *ItemController) GetItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	item, err := ctrl.Service.GetItemByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// UpdateItem updates an item by ID
func (ctrl *ItemController) UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = uint(id) // Set the ID from the URL parameter

	if err := ctrl.Service.UpdateItem(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// DeleteItem deletes an item by ID
func (ctrl *ItemController) DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.Service.DeleteItem(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil) // 204 No Content
}
