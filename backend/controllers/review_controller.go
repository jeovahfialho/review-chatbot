package controllers

import (
	"net/http"
	"review-chatbot/models"
	"review-chatbot/services"
	"time"

	"github.com/gin-gonic/gin"
)

type ReviewController struct {
	service services.ReviewService
}

func NewReviewController(service services.ReviewService) *ReviewController {
	return &ReviewController{service: service}
}

func (rc *ReviewController) CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review.ReviewTime = time.Now() // Adicionar a hora da revisão

	if err := rc.service.CreateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

func (rc *ReviewController) GetReviews(c *gin.Context) {
	reviews, err := rc.service.GetReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
