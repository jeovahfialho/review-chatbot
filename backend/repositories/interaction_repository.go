package repositories

import (
	"errors"
	"log"
	"review-chatbot/models"

	"github.com/jmoiron/sqlx"
)

type InteractionRepository interface {
	CreateInteraction(interaction *models.Interaction) error
}

type interactionRepository struct {
	db           *sqlx.DB
	customerRepo CustomerRepository
}

func NewInteractionRepository(db *sqlx.DB, customerRepo CustomerRepository) InteractionRepository {
	return &interactionRepository{db, customerRepo}
}

func (r *interactionRepository) CreateInteraction(interaction *models.Interaction) error {
	exists, err := r.customerRepo.CustomerExists(interaction.CustomerID)
	if err != nil {
		log.Printf("Error checking if customer exists: %v", err)
		return err
	}
	if !exists {
		log.Printf("Customer does not exist: %d", interaction.CustomerID)
		return errors.New("customer does not exist")
	}
	query := "INSERT INTO `interactions` (`customer_id`, `interaction_time`, `message`) VALUES (?, ?, ?)"
	_, err = r.db.Exec(query, interaction.CustomerID, interaction.InteractionTime, interaction.Message)
	if err != nil {
		log.Printf("Error inserting interaction: %v", err)
	}
	return err
}
