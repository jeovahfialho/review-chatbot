package repositories

import (
	"errors"
	"log"
	"review-chatbot/models"
	"time"

	"github.com/jmoiron/sqlx"
)

// ReviewRepository defines the methods for interacting with review data
type ReviewRepository interface {
	CreateReview(review *models.Review) error                                                                        // Create a new review
	GetReviews() ([]models.Review, error)                                                                            // Retrieve all reviews
	GetAverageRating(startDate, endDate time.Time) (float64, error)                                                  // Get average rating of reviews within a date range
	GetAverageRatingsInIntervals(startDate, endDate time.Time, intervalMinutes int) ([]models.IntervalRating, error) // Get average ratings in intervals
}

// reviewRepository is the implementation of the ReviewRepository interface
type reviewRepository struct {
	db           *sqlx.DB
	customerRepo CustomerRepository
}

// NewReviewRepository creates a new instance of reviewRepository
func NewReviewRepository(db *sqlx.DB, customerRepo CustomerRepository) ReviewRepository {
	return &reviewRepository{db, customerRepo}
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

// GetAverageRating calculates and returns the average rating of reviews within a date range
func (r *reviewRepository) GetAverageRating(startDate, endDate time.Time) (float64, error) {
	var avgRating float64
	query := "SELECT COALESCE(AVG(rating), 0) as avg_rating FROM `reviews` WHERE `review_time` BETWEEN ? AND ?"
	err := r.db.Get(&avgRating, query, startDate, endDate)
	if err != nil {
		return 0, err
	}
	return avgRating, nil
}

// GetAverageRatingsInIntervals calculates the average ratings of reviews in intervals within a date range
func (r *reviewRepository) GetAverageRatingsInIntervals(startDate, endDate time.Time, intervalMinutes int) ([]models.IntervalRating, error) {
	var intervalRatings []models.IntervalRating

	for currentStart := startDate; currentStart.Before(endDate); currentStart = currentStart.Add(time.Duration(intervalMinutes) * time.Minute) {
		currentEnd := currentStart.Add(time.Duration(intervalMinutes) * time.Minute)
		if currentEnd.After(endDate) {
			currentEnd = endDate
		}

		var avgRating float64
		query := "SELECT COALESCE(AVG(rating), 0) as avg_rating FROM `reviews` WHERE `review_time` BETWEEN ? AND ?"
		err := r.db.Get(&avgRating, query, currentStart, currentEnd)
		if err != nil {
			return nil, err
		}

		intervalRating := models.IntervalRating{
			StartDate:     currentStart,
			EndDate:       currentEnd,
			AverageRating: avgRating,
		}
		intervalRatings = append(intervalRatings, intervalRating)
	}

	return intervalRatings, nil
}
