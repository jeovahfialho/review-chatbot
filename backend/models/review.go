package models

import "time"

type Review struct {
	ID         int       `db:"id"`
	CustomerID int       `db:"customer_id"`
	ProductID  int       `db:"product_id"`
	Rating     int       `db:"rating"`
	Comments   string    `db:"comments"`
	ReviewTime time.Time `db:"review_time"`
}
