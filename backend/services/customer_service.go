package services

import (
	"review-chatbot/models"
	"review-chatbot/repositories"
)

type CustomerService interface {
	CreateCustomer(customer *models.Customer) (int, error)
	GetCustomers() ([]models.Customer, error)
	GetCustomerByID(id int) (*models.Customer, error)
}

type customerService struct {
	repo repositories.CustomerRepository
}

func NewCustomerService(repo repositories.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) CreateCustomer(customer *models.Customer) (int, error) {
	return s.repo.CreateCustomer(customer)
}

func (s *customerService) GetCustomers() ([]models.Customer, error) {
	return s.repo.GetCustomers()
}

func (s *customerService) GetCustomerByID(id int) (*models.Customer, error) {
	return s.repo.GetCustomerByID(id)
}
