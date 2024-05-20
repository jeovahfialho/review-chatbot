package services

import (
	"review-chatbot/models"
	"review-chatbot/repositories"
)

// CustomerService defines the methods for customer-related operations
type CustomerService interface {
	CreateCustomer(customer *models.Customer) (int, error) // Create a new customer
	GetCustomers() ([]models.Customer, error)              // Retrieve all customers
	GetCustomerByID(id int) (*models.Customer, error)      // Retrieve a customer by their ID
}

// customerService is the implementation of the CustomerService interface
type customerService struct {
	repo repositories.CustomerRepository
}

// NewCustomerService creates a new instance of customerService
func NewCustomerService(repo repositories.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

// CreateCustomer creates a new customer
func (s *customerService) CreateCustomer(customer *models.Customer) (int, error) {
	return s.repo.CreateCustomer(customer) // Delegate to the repository method
}

// GetCustomers retrieves all customers
func (s *customerService) GetCustomers() ([]models.Customer, error) {
	return s.repo.GetCustomers() // Delegate to the repository method
}

// GetCustomerByID retrieves a customer by their ID
func (s *customerService) GetCustomerByID(id int) (*models.Customer, error) {
	return s.repo.GetCustomerByID(id) // Delegate to the repository method
}
