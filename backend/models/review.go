package models

import "time"

// Review represents a review left by a customer
type Review struct {
	ID         int       `db:"id"`
	CustomerID int       `db:"customer_id"`
	ProductID  int       `db:"product_id"`
	Rating     int       `db:"rating"`
	Comments   string    `db:"comments"`
	ReviewTime time.Time `db:"review_time"`
}

// IntervalRating represents the average rating of reviews within a specific interval
type IntervalRating struct {
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	AverageRating float64   `json:"average_rating"`
}
