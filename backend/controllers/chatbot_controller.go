package controllers

import (
	"net/http"
	"review-chatbot/models"
	"review-chatbot/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatbotHandler struct {
	interactionService services.InteractionService
	reviewService      services.ReviewService
}

func NewChatbotHandler(interactionService services.InteractionService, reviewService services.ReviewService) *ChatbotHandler {
	return &ChatbotHandler{
		interactionService: interactionService,
		reviewService:      reviewService,
	}
}

func (h *ChatbotHandler) Handle(c *gin.Context) {
	var req struct {
		CustomerID int    `json:"customer_id"`
		Message    string `json:"message"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simulando o estado da conversa
	session := h.getSession(req.CustomerID)
	switch session.State {
	case "initial":
		session.State = "ask_rating"
		h.setSession(req.CustomerID, session)
		c.JSON(http.StatusOK, gin.H{"response": "Great! On a scale of 1-5, how would you rate the iPhone 13?"})
	case "ask_rating":
		rating, err := strconv.Atoi(req.Message)
		if err != nil || rating < 1 || rating > 5 {
			c.JSON(http.StatusBadRequest, gin.H{"response": "Please provide a rating between 1 and 5."})
			return
		}
		session.Rating = rating
		session.State = "ask_comment"
		h.setSession(req.CustomerID, session)
		c.JSON(http.StatusOK, gin.H{"response": "Great! Please provide any additional comments about the product."})
	case "ask_comment":
		session.Comment = req.Message
		session.State = "finished"
		h.setSession(req.CustomerID, session)

		review := &models.Review{
			CustomerID: req.CustomerID,
			ProductID:  1, // Substitua pelo ID real do produto
			Rating:     session.Rating,
			Comments:   session.Comment,
			ReviewTime: time.Now(),
		}

		if err := h.reviewService.CreateReview(review); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": "Thank you for your feedback! If you have any more thoughts or need assistance with anything else, feel free to reach out!"})
	default:
		c.JSON(http.StatusOK, gin.H{"response": "Thank you for your message. How can I assist you further?"})
	}
}

// Simulação de uma sessão em memória (substitua por uma solução mais robusta conforme necessário)
var sessions = make(map[int]*Session)

type Session struct {
	State   string
	Rating  int
	Comment string
}

func (h *ChatbotHandler) getSession(customerID int) *Session {
	session, exists := sessions[customerID]
	if !exists {
		session = &Session{State: "initial"}
		sessions[customerID] = session
	}
	return session
}

func (h *ChatbotHandler) setSession(customerID int, session *Session) {
	sessions[customerID] = session
}
