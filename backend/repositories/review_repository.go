package repositories

import (
	"errors"
	"log"
	"review-chatbot/models"

	"github.com/jmoiron/sqlx"
)

// ReviewRepository defines the methods for interacting with review data
type ReviewRepository interface {
	CreateReview(review *models.Review) error // Create a new review
	GetReviews() ([]models.Review, error)     // Retrieve all reviews
}

// reviewRepository is the implementation of the ReviewRepository interface
type reviewRepository struct {
	db           *sqlx.DB
	customerRepo CustomerRepository
}

// NewReviewRepository creates a new instance of reviewRepository
func NewReviewRepository(db *sqlx.DB, customerRepo CustomerRepository) ReviewRepository {
	return &reviewRepository{db: db, customerRepo: customerRepo}
}

// GetReviews retrieves all reviews
func (r *reviewRepository) GetReviews() ([]models.Review, error) {
	var reviews []models.Review
	query := "SELECT `id`, `customer_id`, `product_id`, `rating`, `comments`, `review_time` FROM `reviews`"
	err := r.db.Select(&reviews, query)
	if err != nil {
		return nil, err // Return an error if the query fails
	}
	return reviews, nil // Return the list of reviews
}

// CreateReview creates a new review
func (r *reviewRepository) CreateReview(review *models.Review) error {
	// Check if the customer exists
	exists, err := r.customerRepo.CustomerExists(review.CustomerID)
	if err != nil {
		log.Printf("Error checking if customer exists: %v", err)
		return err
	}
	if !exists {
		log.Printf("Customer does not exist: %d", review.CustomerID)
		return errors.New("customer does not exist")
	}

	// Insert the review into the database
	query := "INSERT INTO `reviews` (`customer_id`, `product_id`, `rating`, `comments`, `review_time`) VALUES (?, ?, ?, ?, ?)"
	_, err = r.db.Exec(query, review.CustomerID, review.ProductID, review.Rating, review.Comments, review.ReviewTime)
	if err != nil {
		log.Printf("Error inserting review: %v", err)
	}
	return err // Return any error that occurred during the insertion
}
