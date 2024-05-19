package controllers

import (
	"log"
	"net/http"
	"os"
	"review-chatbot/models"
	"review-chatbot/services"
	"review-chatbot/workflow"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatbotHandler struct {
	interactionService services.InteractionService
	reviewService      services.ReviewService
	orderService       services.OrderService
	workflow           workflow.Workflow
}

func NewChatbotHandler(interactionService services.InteractionService, reviewService services.ReviewService, orderService services.OrderService) *ChatbotHandler {
	workflowPaths := []string{
		"./controllers/workflow.json",
		"./workflow.json",
	}

	var loadedWorkflow workflow.Workflow
	var err error

	for _, path := range workflowPaths {
		if _, err := os.Stat(path); err == nil {
			loadedWorkflow, err = workflow.LoadWorkflow(path)
			if err == nil {
				break
			}
		}
	}

	if err != nil {
		panic("Failed to load workflow.json from any known path")
	}

	return &ChatbotHandler{
		interactionService: interactionService,
		reviewService:      reviewService,
		orderService:       orderService,
		workflow:           loadedWorkflow,
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

	session := h.getSession(req.CustomerID)

	// Sempre começar com a mensagem de saudação se o estado da sessão for "greeting"
	if session.State == "greeting" && (req.Message == "hi" || req.Message == "hello") {
		currentStep := h.workflow["default"]["greeting"]
		session.State = "options"
		h.setSession(req.CustomerID, session)
		log.Printf("Greeting Step: %+v", currentStep)
		log.Printf("Greeting Step Response: %s", currentStep.Response)
		c.JSON(http.StatusOK, gin.H{"response": "Sales: " + currentStep.Response})
		return
	}

	currentStep, nextState := workflow.GetNextStep(h.workflow, session.State, req.Message)

	log.Printf("Current Step: %+v", currentStep)
	log.Printf("Current Step Response: %s", currentStep.Response)

	if currentStep.Action != "" {
		// Perform action if any
		switch currentStep.Action {
		case "getOrderStatus":
			orderID, err := strconv.Atoi(req.Message)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
				return
			}
			status, err := h.orderService.GetOrderStatus(orderID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			currentStep.Response = status
		case "initiateReturn":
			orderID, err := strconv.Atoi(req.Message)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
				return
			}
			status, err := h.orderService.InitiateReturn(orderID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			currentStep.Response = status
		case "saveReview":
			err := h.saveReview(session, req.Message)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	session.State = nextState
	h.setSession(req.CustomerID, session)
	responseMessage := "Sales: " + currentStep.Response
	log.Printf("Response Message: %s", responseMessage)
	c.JSON(http.StatusOK, gin.H{"response": responseMessage})
}

func (h *ChatbotHandler) saveReview(session *Session, comments string) error {
	rating, _ := strconv.Atoi(session.Data["rating"])
	review := &models.Review{
		CustomerID: session.CustomerID,
		ProductID:  1, // Substitute with actual product ID
		Rating:     rating,
		Comments:   comments,
		ReviewTime: time.Now(),
	}
	return h.reviewService.CreateReview(review)
}

// Session management code remains the same
type Session struct {
	CustomerID int
	State      string
	Data       map[string]string
}

var sessions = make(map[int]*Session)

func (h *ChatbotHandler) getSession(customerID int) *Session {
	session, exists := sessions[customerID]
	if !exists {
		session = &Session{
			CustomerID: customerID,
			State:      "greeting",
			Data:       make(map[string]string),
		}
		sessions[customerID] = session
	}
	return session
}

func (h *ChatbotHandler) setSession(customerID int, session *Session) {
	sessions[customerID] = session
}
