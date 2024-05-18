package services

import (
	"review-chatbot/models"
	"review-chatbot/repositories"
	"time"
)

type InteractionService interface {
	CreateInteraction(customerID int, message string) error
}

type interactionService struct {
	repo repositories.InteractionRepository
}

func NewInteractionService(repo repositories.InteractionRepository) InteractionService {
	return &interactionService{repo}
}

func (s *interactionService) CreateInteraction(customerID int, message string) error {
	interaction := &models.Interaction{
		CustomerID:      customerID,
		InteractionTime: time.Now(),
		Message:         message,
	}
	return s.repo.CreateInteraction(interaction)
}
