package services

import (
	"review-chatbot/models"
	"review-chatbot/repositories"
	"time"
)

// InteractionService defines the methods for interaction-related operations
type InteractionService interface {
	CreateInteraction(customerID int, message string) error // Create a new interaction
}

// interactionService is the implementation of the InteractionService interface
type interactionService struct {
	repo repositories.InteractionRepository
}

// NewInteractionService creates a new instance of interactionService
func NewInteractionService(repo repositories.InteractionRepository) InteractionService {
	return &interactionService{repo}
}

// CreateInteraction creates a new interaction
func (s *interactionService) CreateInteraction(customerID int, message string) error {
	// Create a new interaction model with the provided customer ID, current time, and message
	interaction := &models.Interaction{
		CustomerID:      customerID,
		InteractionTime: time.Now(),
		Message:         message,
	}
	// Delegate to the repository method to create the interaction
	return s.repo.CreateInteraction(interaction)
}
