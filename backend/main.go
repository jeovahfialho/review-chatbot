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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With, Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

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

	router := gin.Default()

	// Use the CORS middleware
	router.Use(CORSMiddleware())

	// API routes
	api := router.Group("/api")
	{
		api.POST("/customers", customerController.CreateCustomer)
		api.GET("/customers", customerController.GetCustomers)
		api.GET("/customers/:id", customerController.GetCustomerByID)
		api.POST("/interaction", interactionController.CreateInteraction)
		api.POST("/review", reviewController.CreateReview)
		api.GET("/reviews", reviewController.GetReviews)
		api.POST("/chatbot", chatbotHandler.Handle)
	}

	log.Println("Serving APIs on port 8080...")
	log.Fatal(router.Run(":8080"))
}
