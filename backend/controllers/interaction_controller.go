package controllers

import (
	"net/http"
	"review-chatbot/services"

	"github.com/gin-gonic/gin"
)

type InteractionController struct {
	service services.InteractionService
}

func NewInteractionController(service services.InteractionService) *InteractionController {
	return &InteractionController{service}
}

func (c *InteractionController) CreateInteraction(ctx *gin.Context) {
	var request struct {
		CustomerID int    `json:"customer_id"`
		Message    string `json:"message"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.CreateInteraction(request.CustomerID, request.Message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "interaction recorded"})
}
