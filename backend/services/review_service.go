package services

import (
	"review-chatbot/models"
	"review-chatbot/repositories"
)

type ReviewService interface {
	CreateReview(review *models.Review) error
	GetReviews() ([]models.Review, error)
}

type reviewService struct {
	repo repositories.ReviewRepository
}

func NewReviewService(repo repositories.ReviewRepository) ReviewService {
	return &reviewService{repo: repo}
}

func (s *reviewService) CreateReview(review *models.Review) error {
	return s.repo.CreateReview(review)
}

func (s *reviewService) GetReviews() ([]models.Review, error) {
	return s.repo.GetReviews()
}
