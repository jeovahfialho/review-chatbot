package controllers

import (
	"log"
	"net/http"
	"review-chatbot/models"
	"review-chatbot/services"
	"time"

	"github.com/gin-gonic/gin"
)

// ReviewController handles review-related requests
type ReviewController struct {
	service services.ReviewService
}

// NewReviewController creates a new ReviewController
func NewReviewController(service services.ReviewService) *ReviewController {
	return &ReviewController{service: service}
}

// CreateReview handles the creation of a new review
func (rc *ReviewController) CreateReview(c *gin.Context) {
	var req ReviewRequest
	// Bind the JSON body of the request to the ReviewRequest structure
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Map the ReviewRequest to the Review model
	review := models.Review{
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Rating:     req.Rating,
		Comments:   req.Comments,
		ReviewTime: time.Now(), // Add the current time to the review
	}

	// Print the raw JSON request
	log.Printf("POST /api/review - Raw JSON: %+v", review)

	// Create a new review using the service
	if err := rc.service.CreateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created review
	c.JSON(http.StatusOK, review)
}

// GetReviews handles the retrieval of all reviews
func (rc *ReviewController) GetReviews(c *gin.Context) {
	// Get the list of reviews from the service
	reviews, err := rc.service.GetReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of reviews
	c.JSON(http.StatusOK, reviews)
}
