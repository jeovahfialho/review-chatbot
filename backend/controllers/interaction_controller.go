package controllers

import (
	"net/http"
	"review-chatbot/services"

	"github.com/gin-gonic/gin"
)

// InteractionController handles interactions with the chatbot
type InteractionController struct {
	service services.InteractionService
}

// NewInteractionController creates a new instance of InteractionController
func NewInteractionController(service services.InteractionService) *InteractionController {
	return &InteractionController{service}
}

// CreateInteraction handles the creation of a new interaction
func (c *InteractionController) CreateInteraction(ctx *gin.Context) {
	// Define the structure of the incoming request
	var request struct {
		CustomerID int    `json:"customer_id"`
		Message    string `json:"message"`
	}

	// Bind the JSON body of the request to the request structure
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new interaction using the service
	err := c.service.CreateInteraction(request.CustomerID, request.Message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success status
	ctx.JSON(http.StatusOK, gin.H{"status": "interaction recorded"})
}
