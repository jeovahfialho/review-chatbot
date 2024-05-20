package controllers

import (
	"net/http"
	"review-chatbot/models"
	"review-chatbot/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CustomerController handles customer-related operations
type CustomerController struct {
	service services.CustomerService
}

// NewCustomerController creates a new instance of CustomerController
func NewCustomerController(service services.CustomerService) *CustomerController {
	return &CustomerController{service: service}
}

// CreateCustomer handles the creation of a new customer
func (cc *CustomerController) CreateCustomer(c *gin.Context) {
	var customer models.Customer
	// Bind the JSON body of the request to the customer structure
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new customer using the service
	id, err := cc.service.CreateCustomer(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set the ID of the created customer and return the customer object
	customer.ID = id
	c.JSON(http.StatusOK, customer)
}

// GetCustomers handles the retrieval of all customers
func (cc *CustomerController) GetCustomers(c *gin.Context) {
	// Get the list of customers from the service
	customers, err := cc.service.GetCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of customers
	c.JSON(http.StatusOK, customers)
}

// GetCustomerByID handles the retrieval of a customer by their ID
func (cc *CustomerController) GetCustomerByID(c *gin.Context) {
	// Get the customer ID from the request parameters and convert it to an integer
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	// Get the customer by ID using the service
	customer, err := cc.service.GetCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the customer object
	c.JSON(http.StatusOK, customer)
}
