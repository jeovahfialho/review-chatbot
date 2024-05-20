package repositories

import (
	"review-chatbot/models"

	"github.com/jmoiron/sqlx"
)

// CustomerRepository defines the methods for interacting with the customer data
type CustomerRepository interface {
	GetCustomerByID(id int) (*models.Customer, error)      // Retrieve a customer by their ID
	CustomerExists(id int) (bool, error)                   // Check if a customer exists by their ID
	CreateCustomer(customer *models.Customer) (int, error) // Create a new customer
	GetCustomers() ([]models.Customer, error)              // Retrieve all customers
}

// customerRepository is the implementation of the CustomerRepository interface
type customerRepository struct {
	db *sqlx.DB
}

// NewCustomerRepository creates a new instance of customerRepository
func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return &customerRepository{db}
}

// GetCustomerByID retrieves a customer by their ID
func (r *customerRepository) GetCustomerByID(id int) (*models.Customer, error) {
	var customer models.Customer
	query := "SELECT * FROM customers WHERE id = ?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err // Return an error if the query fails
	}
	return &customer, nil // Return the customer
}

// CustomerExists checks if a customer exists by their ID
func (r *customerRepository) CustomerExists(id int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM customers WHERE id = ?)"
	err := r.db.Get(&exists, query, id)
	return exists, err // Return the existence check result and any error
}

// CreateCustomer creates a new customer
func (r *customerRepository) CreateCustomer(customer *models.Customer) (int, error) {
	query := "INSERT INTO customers (name, email) VALUES (?, ?)"
	result, err := r.db.Exec(query, customer.Name, customer.Email)
	if err != nil {
		return 0, err // Return an error if the query fails
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err // Return an error if getting the last insert ID fails
	}
	return int(id), nil // Return the new customer ID
}

// GetCustomers retrieves all customers
func (r *customerRepository) GetCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	query := "SELECT * FROM customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err // Return an error if the query fails
	}
	return customers, nil // Return the list of customers
}
