package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"review-chatbot/models"
	"review-chatbot/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ChatbotHandler handles chatbot interactions
type ChatbotHandler struct {
	interactionService services.InteractionService
	reviewService      services.ReviewService
}

// NewChatbotHandler creates a new ChatbotHandler
func NewChatbotHandler(interactionService services.InteractionService, reviewService services.ReviewService) *ChatbotHandler {
	return &ChatbotHandler{
		interactionService: interactionService,
		reviewService:      reviewService,
	}
}

// Handle processes incoming chatbot messages
func (h *ChatbotHandler) Handle(c *gin.Context) {
	var req struct {
		CustomerID int    `json:"customer_id"`
		Message    string `json:"message"`
	}

	// Bind JSON payload to the request structure
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the session for the customer
	session := h.getSession(req.CustomerID)
	switch session.State {
	case "initial":
		session.State = "ask_rating"
		h.setSession(req.CustomerID, session)
		c.JSON(http.StatusOK, gin.H{"response": "Great! On a scale of 1-5, how would you rate the product?"})
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

		// Create a new review with the provided data
		review := &models.Review{
			CustomerID: req.CustomerID,
			ProductID:  1, // Replace with the actual product ID
			Rating:     session.Rating,
			Comments:   session.Comment,
			ReviewTime: time.Now(),
		}

		// Convert the review to JSON
		reviewJSON, err := json.Marshal(review)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review JSON"})
			return
		}

		// Log the raw JSON request
		log.Printf("POST /api/review - Raw JSON: %s", string(reviewJSON))

		if err := h.reviewService.CreateReview(review); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": "Thank you for your feedback! If you need anything else, I'm here to help."})
	case "ask_return":
		session.State = "return_initiated"
		h.setSession(req.CustomerID, session)
		c.JSON(http.StatusOK, gin.H{"response": "Thank you. Your return request has been initiated. You will receive further instructions via email."})
	case "ask_recommendation":
		session.State = "recommendation_given"
		h.setSession(req.CustomerID, session)
		c.JSON(http.StatusOK, gin.H{"response": "Here are some product recommendations: [product recommendations]."})
	default:
		if req.Message == "I want to return a product" {
			session.State = "ask_return"
			h.setSession(req.CustomerID, session)
			c.JSON(http.StatusOK, gin.H{"response": "I'm sorry to hear that. Can you please provide the order number?"})
		} else if req.Message == "Can you recommend a product?" {
			session.State = "ask_recommendation"
			h.setSession(req.CustomerID, session)
			c.JSON(http.StatusOK, gin.H{"response": "Sure! What kind of product are you looking for?"})
		} else {
			c.JSON(http.StatusOK, gin.H{"response": "Thank you for your message. How can I assist you further?"})
		}
	}
}

// sessions stores the state of each customer's session
var sessions = make(map[int]*Session)

// Session represents the state of a conversation with a customer
type Session struct {
	State   string
	Rating  int
	Comment string
}

// getSession retrieves the session for a given customer, creating a new one if necessary
func (h *ChatbotHandler) getSession(customerID int) *Session {
	session, exists := sessions[customerID]
	if !exists {
		session = &Session{State: "initial"}
		sessions[customerID] = session
	}
	return session
}

// setSession saves the session for a given customer
func (h *ChatbotHandler) setSession(customerID int, session *Session) {
	sessions[customerID] = session
}
