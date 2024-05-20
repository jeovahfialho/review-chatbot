package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"review-chatbot/models"
	"review-chatbot/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SentimentResponse represents a single response item from the sentiment analysis service
type SentimentResponse struct {
	Label string  `json:"label"`
	Score float64 `json:"score"`
}

// SentimentRequest represents the request payload for the sentiment analysis service
type SentimentRequest struct {
	Text string `json:"text"`
}

// ReviewRequest is used to bind the incoming JSON payload
type ReviewRequest struct {
	CustomerID int    `json:"customer_id"`
	ProductID  int    `json:"product_id"`
	Rating     int    `json:"rating"`
	Comments   string `json:"comments"`
}

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

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := h.getSession(req.CustomerID)
	switch session.State {
	case "initial":
		if req.Message == "start_conversation" {
			session.State = "ask_rating"
			h.setSession(req.CustomerID, session)
			c.JSON(http.StatusOK, gin.H{"response": "Hello! We noticed you've recently received your product. We'd love to hear about your experience. Can you share your thoughts?"})
		} else if req.Message == "Hi, I am here again." {
			c.JSON(http.StatusOK, gin.H{"response": "Thank you for your message. How can I assist you further?"})
		} else {
			c.JSON(http.StatusOK, gin.H{"response": "Thank you for your message. How can I assist you further?"})
		}
	case "ask_rating":
		if req.Message == "Yes" {
			session.State = "ask_rating_value"
			h.setSession(req.CustomerID, session)
			c.JSON(http.StatusOK, gin.H{"response": "Great! On a scale of 1 to 5, how would you rate the product?"})
		} else {
			c.JSON(http.StatusOK, gin.H{"response": "Can you please confirm if you'd like to share your thoughts on the product?"})
		}
	case "ask_rating_value":
		rating, err := strconv.Atoi(req.Message)
		if err != nil || rating < 1 || rating > 5 {
			c.JSON(http.StatusBadRequest, gin.H{"response": "Please provide a rating between 1 and 5."})
			return
		}
		session.Rating = rating
		session.State = "ask_comment"
		h.setSession(req.CustomerID, session)
		c.JSON(http.StatusOK, gin.H{"response": "Thank you for your rating! Would you like to leave any additional comments about the product?"})
	case "ask_comment":
		if req.Message != "" {
			session.Comment = req.Message
			session.State = "finished"
			h.setSession(req.CustomerID, session)

			review := &models.Review{
				CustomerID: req.CustomerID,
				ProductID:  1, // Replace with the actual product ID
				Rating:     session.Rating,
				Comments:   session.Comment,
				ReviewTime: time.Now(),
				// Add field to store sentiment analysis result if needed
			}

			sentiment, err := analyzeSentiment(review.Comments)
			if err != nil {
				log.Printf("Error analyzing sentiment: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to analyze sentiment"})
				return
			}

			log.Printf("Sentiment analysis result: %+v", sentiment)

			log.Printf("POST /api/review - Raw JSON: %+v", review)

			if err := h.reviewService.CreateReview(review); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"response": "Thank you for your feedback! If you need anything else, I'm here to help."})
		} else {
			c.JSON(http.StatusOK, gin.H{"response": "Can you please confirm your comment on the product?"})
		}
	case "finished":
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
	case "ask_return":
		session.State = "return_initiated"
		h.setSession(req.CustomerID, session)
		c.JSON(http.StatusOK, gin.H{"response": "Thank you. Your return request has been initiated. You will receive further instructions via email."})
	case "ask_recommendation":
		session.State = "recommendation_given"
		h.setSession(req.CustomerID, session)
		c.JSON(http.StatusOK, gin.H{"response": "Here are some recommendations: [product recommendations]."})
	default:
		c.JSON(http.StatusOK, gin.H{"response": "Thank you for your message. How can I assist you further?"})
	}
}

// analyzeSentiment calls the sentiment analysis API
func analyzeSentiment(text string) ([]SentimentResponse, error) {
	sentimentReq := SentimentRequest{Text: text}
	jsonValue, _ := json.Marshal(sentimentReq)
	resp, err := http.Post("http://sentiment-analysis:5000/sentiment", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var sentimentResp []SentimentResponse
	if err := json.NewDecoder(resp.Body).Decode(&sentimentResp); err != nil {
		return nil, err
	}

	return sentimentResp, nil
}

var sessions = make(map[int]*Session)

// Session represents the state of a conversation with a customer
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
