package controllers

import (
	"net/http"
	"review-chatbot/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ReviewStatsController struct {
	service services.ReviewService
}

func NewReviewStatsController(service services.ReviewService) *ReviewStatsController {
	return &ReviewStatsController{service}
}

// GetAverageRating handles the request to get the average rating of reviews within a date range
func (rc *ReviewStatsController) GetAverageRating(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	intervalStr := c.Query("interval_minutes")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format. Use YYYY-MM-DD."})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format. Use YYYY-MM-DD."})
		return
	}

	intervalMinutes, err := strconv.Atoi(intervalStr)
	if err != nil || intervalMinutes <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval. Must be a positive integer representing minutes."})
		return
	}

	avgRating, err := rc.service.GetAverageRating(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	intervalRatings, err := rc.service.GetAverageRatingsInIntervals(startDate, endDate, intervalMinutes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"start_date":           startDateStr,
		"end_date":             endDateStr,
		"average_rating_total": avgRating,
		"interval_ratings":     intervalRatings,
	})
}
