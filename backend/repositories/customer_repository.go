package repositories

import (
	"review-chatbot/models"

	"github.com/jmoiron/sqlx"
)

type CustomerRepository interface {
	GetCustomerByID(id int) (*models.Customer, error)
	CustomerExists(id int) (bool, error)
	CreateCustomer(customer *models.Customer) (int, error)
	GetCustomers() ([]models.Customer, error)
}

type customerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) GetCustomerByID(id int) (*models.Customer, error) {
	var customer models.Customer
	query := "SELECT * FROM customers WHERE id = ?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) CustomerExists(id int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM customers WHERE id = ?)"
	err := r.db.Get(&exists, query, id)
	return exists, err
}

func (r *customerRepository) CreateCustomer(customer *models.Customer) (int, error) {
	query := "INSERT INTO customers (name, email) VALUES (?, ?)"
	result, err := r.db.Exec(query, customer.Name, customer.Email)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *customerRepository) GetCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	query := "SELECT * FROM customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}
