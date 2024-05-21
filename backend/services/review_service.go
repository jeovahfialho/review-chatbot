package services

import (
	"review-chatbot/models"
	"review-chatbot/repositories"
	"time"
)

// ReviewService defines the methods for review-related operations
type ReviewService interface {
	CreateReview(review *models.Review) error                       // Create a new review
	GetReviews() ([]models.Review, error)                           // Retrieve all reviews
	GetAverageRating(startDate, endDate time.Time) (float64, error) // Get Average rating of reviews
}

// reviewService is the implementation of the ReviewService interface
type reviewService struct {
	repo repositories.ReviewRepository
}

// NewReviewService creates a new instance of reviewService
func NewReviewService(repo repositories.ReviewRepository) ReviewService {
	return &reviewService{repo: repo}
}

// CreateReview creates a new review
func (s *reviewService) CreateReview(review *models.Review) error {
	return s.repo.CreateReview(review) // Delegate to the repository method
}

// GetReviews retrieves all reviews
func (s *reviewService) GetReviews() ([]models.Review, error) {
	return s.repo.GetReviews() // Delegate to the repository method
}

func (s *reviewService) GetAverageRating(startDate, endDate time.Time) (float64, error) {
	return s.repo.GetAverageRating()
}
