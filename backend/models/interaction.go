package models

import "time"

type Interaction struct {
	ID              int       `db:"id"`
	CustomerID      int       `db:"customer_id"`
	InteractionTime time.Time `db:"interaction_time"`
	Message         string    `db:"message"`
}
