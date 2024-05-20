package main

import (
	"log"
	"net/http"
	"review-chatbot/controllers"
	"review-chatbot/database"
	"review-chatbot/repositories"
	"review-chatbot/services"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware handles Cross-Origin Resource Sharing (CORS) settings
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the headers for CORS
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With, Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repositories
	customerRepo := repositories.NewCustomerRepository(db)
	reviewRepo := repositories.NewReviewRepository(db, customerRepo)
	interactionRepo := repositories.NewInteractionRepository(db, customerRepo)

	// Initialize services
	customerService := services.NewCustomerService(customerRepo)
	reviewService := services.NewReviewService(reviewRepo)
	interactionService := services.NewInteractionService(interactionRepo)

	// Initialize controllers with the services
	customerController := controllers.NewCustomerController(customerService)
	reviewController := controllers.NewReviewController(reviewService)
	interactionController := controllers.NewInteractionController(interactionService)
	chatbotHandler := controllers.NewChatbotHandler(interactionService, reviewService)

	// Create a new Gin router
	router := gin.Default()

	// Use the CORS middleware
	router.Use(CORSMiddleware())

	// API routes
	api := router.Group("/api")
	{
		api.POST("/customers", customerController.CreateCustomer)         // Create a new customer
		api.GET("/customers", customerController.GetCustomers)            // Get all customers
		api.GET("/customers/:id", customerController.GetCustomerByID)     // Get a customer by ID
		api.POST("/interaction", interactionController.CreateInteraction) // Create a new interaction
		api.POST("/review", reviewController.CreateReview)                // Create a new review
		api.GET("/reviews", reviewController.GetReviews)                  // Get all reviews
		api.POST("/chatbot", chatbotHandler.Handle)                       // Handle chatbot interactions
	}

	// Start the server on port 8080
	log.Println("Serving APIs on port 8080...")
	log.Fatal(router.Run(":8080"))
}
