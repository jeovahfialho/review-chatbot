package repositories

import (
	"errors"
	"log"
	"review-chatbot/models"

	"github.com/jmoiron/sqlx"
)

type ReviewRepository interface {
	CreateReview(review *models.Review) error
	GetReviews() ([]models.Review, error)
}

type reviewRepository struct {
	db           *sqlx.DB
	customerRepo CustomerRepository
}

func NewReviewRepository(db *sqlx.DB, customerRepo CustomerRepository) ReviewRepository {
	return &reviewRepository{db, customerRepo}
}

func (r *reviewRepository) GetReviews() ([]models.Review, error) {
	var reviews []models.Review
	query := "SELECT `id`, `customer_id`, `product_id`, `rating`, `comments`, `review_time` FROM `reviews`"
	err := r.db.Select(&reviews, query)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *reviewRepository) CreateReview(review *models.Review) error {
	exists, err := r.customerRepo.CustomerExists(review.CustomerID)
	if err != nil {
		log.Printf("Error checking if customer exists: %v", err)
		return err
	}
	if !exists {
		log.Printf("Customer does not exist: %d", review.CustomerID)
		return errors.New("customer does not exist")
	}
	query := "INSERT INTO `reviews` (`customer_id`, `product_id`, `rating`, `comments`, `review_time`) VALUES (?, ?, ?, ?, ?)"
	_, err = r.db.Exec(query, review.CustomerID, review.ProductID, review.Rating, review.Comments, review.ReviewTime)
	if err != nil {
		log.Printf("Error inserting review: %v", err)
	}
	return err
}
