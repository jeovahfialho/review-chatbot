package controllers

import (
	"net/http"
	"review-chatbot/services"
	"time"

	"github.com/gin-gonic/gin"
)

type ReviewStatsController struct {
	service services.ReviewService
}

func NewReviewStatsController(service services.ReviewService) *ReviewStatsController {
	return &ReviewStatsController{service}
}

func (rc *ReviewStatsController) GetAverageRating(c *gin.Context) {

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, err := time.Parse("2009-01-02", startDateStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format."})
		return
	}

	endDate, err := time.Parse("2009-01-02", endDateStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format."})
		return
	}

	avgRating, err := rc.service.GetAverageRating(startDate, endDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"avarage_rating": avgRating})
}
