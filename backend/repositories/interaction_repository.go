package repositories

import (
	"errors"
	"log"
	"review-chatbot/models"

	"github.com/jmoiron/sqlx"
)

// InteractionRepository defines the methods for interacting with interaction data
type InteractionRepository interface {
	CreateInteraction(interaction *models.Interaction) error // Create a new interaction
}

// interactionRepository is the implementation of the InteractionRepository interface
type interactionRepository struct {
	db           *sqlx.DB
	customerRepo CustomerRepository
}

// NewInteractionRepository creates a new instance of interactionRepository
func NewInteractionRepository(db *sqlx.DB, customerRepo CustomerRepository) InteractionRepository {
	return &interactionRepository{db, customerRepo}
}

// CreateInteraction creates a new interaction
func (r *interactionRepository) CreateInteraction(interaction *models.Interaction) error {
	// Check if the customer exists
	exists, err := r.customerRepo.CustomerExists(interaction.CustomerID)
	if err != nil {
		log.Printf("Error checking if customer exists: %v", err)
		return err
	}
	if !exists {
		log.Printf("Customer does not exist: %d", interaction.CustomerID)
		return errors.New("customer does not exist")
	}

	// Insert the interaction into the database
	query := "INSERT INTO `interactions` (`customer_id`, `interaction_time`, `message`) VALUES (?, ?, ?)"
	_, err = r.db.Exec(query, interaction.CustomerID, interaction.InteractionTime, interaction.Message)
	if err != nil {
		log.Printf("Error inserting interaction: %v", err)
	}
	return err // Return any error that occurred during the insertion
}
