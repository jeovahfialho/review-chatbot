package models

type Customer struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}
