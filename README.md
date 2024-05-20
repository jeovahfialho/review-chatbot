# Review Collection Chatbot

This project implements a chatbot to collect product reviews after delivery confirmation. The chatbot can also engage the user in other relevant conversations, providing an interactive and personalized user experience.

## Conversation Flow

The conversation flow of the review collection chatbot is shown below:


### Conversation Flow Description

1. **Event: The system signals that the delivery has been confirmed.**
   - **Chatbot**: "Hello! We noticed you've recently received [product name]. We'd love to hear about your experience. Can you share your thoughts?"
   - **User**: "Yes"
   - **Chatbot**: "Great! On a scale of 1 to 5, how would you rate [product name]?"
   - **User**: "5"
   - **Chatbot**: "Thank you for your rating! Would you like to leave any additional comments about the product?"
   - **User**: "Yes, it's a great product!"
   - **Chatbot**: "Thank you for your feedback! If you need anything else, I'm here to help."

2. **Event: The user wants to start a new conversation.**
   - **User**: "Hi, I am here again."
   - **Chatbot**: "Thank you for your message. How can I assist you further?"

3. **Event: The user wants to return a product.**
   - **User**: "I want to return a product."
   - **Chatbot**: "I'm sorry to hear that. Can you please provide the order number?"
   - **User**: "[order number]"
   - **Chatbot**: "Thank you. Your return request has been initiated. You will receive further instructions via email."

4. **Event: The user is looking for product recommendations.**
   - **User**: "Can you recommend a product?"
   - **Chatbot**: "Sure! What kind of product are you looking for?"
   - **User**: "[product type]"
   - **Chatbot**: "Here are some recommendations: [product recommendations]."

## Project Structure

The project is divided into the following parts:

### Backend

- **Language**: Go
- **Framework**: Gin
- **Purpose**: Provides the REST API for interacting with the chatbot and handling review data.

#### Directory Structure

```
backend/
├── controllers
│   ├── chatbot_controller.go
│   ├── customer_controller.go
│   ├── interaction_controller.go
│   └── review_controller.go
├── models
│   ├── customer.go
│   ├── interaction.go
│   └── review.go
├── repositories
│   ├── customer_repository.go
│   ├── interaction_repository.go
│   └── review_repository.go
├── services
│   ├── customer_service.go
│   ├── interaction_service.go
│   └── review_service.go
├── main.go
├── go.mod
└── go.sum
```

### Frontend

- **Language**: JavaScript
- **Framework**: Vue.js
- **Purpose**: Provides the user interface for interacting with the chatbot and viewing reviews.

#### Directory Structure

```
frontend/
├── public
│   └── index.html
├── src
│   ├── components
│   │   ├── Chatbot.vue
│   │   ├── CustomerForm.vue
│   │   ├── CustomerTable.vue
│   │   ├── Dashboard.vue
│   │   └── ReviewTable.vue
│   ├── App.vue
│   ├── main.js
│   └── router.js
├── package.json
└── vue.config.js
```

### Database

- **Type**: MySQL/MariaDB
- **Purpose**: Stores customer, interaction, and review data.

#### Schema

```sql
CREATE TABLE IF NOT EXISTS `reviews` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `customer_id` INT NOT NULL,
    `product_id` INT NOT NULL,
    `rating` INT NOT NULL,
    `comments` TEXT,
    `review_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS `customers` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL,
    `email` VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS `interactions` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `customer_id` INT NOT NULL,
    `message` TEXT,
    `timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`customer_id`) REFERENCES `customers`(`id`)
);
```

## Architecture Decisions

- **SOLID Principles**: The project follows SOLID principles to ensure a scalable and maintainable architecture.
- **REST API**: The backend exposes a REST API to handle interactions and reviews.
- **Separation of Concerns**: The project is divided into controllers, services, and repositories to separate business logic, data access, and API handling.
- **Docker**: The project uses Docker and Docker Compose for containerization and orchestration, making it easy to set up and run the project in different environments.

## How to Run the Project

### Prerequisites

- Docker
- Docker Compose

### Steps to Run

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/review-chatbot.git
   cd review-chatbot
   ```

2. Start the containers:

   ```sh
   docker-compose up --build
   ```

3. Access the graphical interface in your browser:

   ```sh
   http://localhost
   ```

### Access the Graphical Interface

- **Customer Registration**: Use the "Register Customer" option to add new customers.
- **Send Message**: Use the "Send Message" option to open the chatbot interface.
- **View Reviews**: The dashboard will display the list of product reviews.

## Manual Testing

### Initial Interaction

1. **Visit the chatbot page**: `http://localhost/chatbot?customerID=1`
2. **Expected Chatbot Message**: "Hello! We noticed you've recently received [product name]. We'd love to hear about your experience. Can you share your thoughts?"
3. **User**: "Yes"
4. **Expected Chatbot Message**: "Great! On a scale of 1 to 5, how would you rate [product name]?"

### Providing a Rating

5. **User**: "5"
6. **Expected Chatbot Message**: "Thank you for your rating! Would you like to leave any additional comments about the product?"

### Providing Additional Comments

7. **User**: "Yes, it's a great product!"
8. **Expected Chatbot Message**: "Thank you for your feedback! If you need anything else, I'm here to help."

### Starting a New Conversation

9. **User**: "Hi, I am here again."
10. **Expected Chatbot Message**: "Thank you for your message. How can I assist you further?"

### Initiating a Product Return

11. **User**: "I want to return a product"
12. **Expected Chatbot Message**: "I'm sorry to hear that. Can you please provide the order number?"
13. **User****: "[order number]"
14. **Expected Chatbot Message**: "Thank you. Your return request has been initiated. You will receive further instructions via email."

### Requesting Product Recommendations

15. **User**: "Can you recommend a product?"
16. **Expected Chatbot Message**: "Sure! What kind of product are you looking for?"
17. **User**: "[product type]"
18. **Expected Chatbot Message**: "Here are some recommendations: [product recommendations]."

### Testing the Endpoints

You can use tools like `curl` or Postman to test the API endpoints.

#### Example with `curl`

- **Create a Customer**:

  ```sh
  curl -X POST http://localhost:8080/api/customers   -H 'Content-Type: application/json'   -d '{"name": "John Doe", "email": "john.doe@example.com"}'
  ```

- **Get All Customers**:

  ```sh
  curl -X GET http://localhost:8080/api/customers
  ```

- **Get a Customer by ID**:

  ```sh
  curl -X GET http://localhost:8080/api/customers/1
  ```

- **Create a Review**:

  ```sh
  curl -X POST http://localhost:8080/api/review   -H 'Content-Type: application/json'   -d '{"customer_id": 1, "product_id": 1, "rating": 5, "comments": "Great product!"}'
  ```

- **Get All Reviews**:

  ```sh
  curl -X GET http://localhost:8080/api/reviews
  ```

- **Chatbot Interaction**:

  ```sh
  curl -X POST http://localhost:8080/api/chatbot   -H 'Content-Type: application/json'   -d '{"customer_id": 1, "message": "Hi"}'
  ```


## Sentiment Analysis Implementation

### Motivation for Sentiment Analysis

The decision to incorporate sentiment analysis into the review flow was motivated by the need to gain deeper insights into customer feedback. While star ratings provide quantitative data, they do not capture the qualitative nuances of customer opinions. Sentiment analysis allows us to understand the emotional tone behind the feedback, which can help in improving products and services more effectively.

### What the Application Does

The application uses a machine learning model to perform sentiment analysis on the comments provided by customers during the review process. When a customer submits a comment, the application sends the text to a sentiment analysis service, which returns the sentiment score and label (e.g., positive, negative, neutral). This information is logged for further analysis and can be used to identify trends in customer feedback.

### Viewing Sentiment Analysis Results

The results of the sentiment analysis are logged in the backend service logs. To view these results, you can check the logs of the backend container. The logs will include entries showing the sentiment analysis result for each comment submitted by customers.

#### Example Log Entry

Here is an example of what the log entries for sentiment analysis results might look like:

```plaintext
2024/05/20 13:45:23 Sentiment analysis result: [{Label: "POSITIVE", Score: 0.95}]
```

### How to Access the Logs

You can access the logs of the backend container using the following Docker command:

```sh
docker-compose logs backend
```

This command will display the logs of the backend service, including the sentiment analysis results.

### Implementation Details

#### Sentiment Analysis Service

The sentiment analysis service is implemented using a Python-based machine learning model. The service is containerized and runs alongside the backend and frontend services. The backend sends HTTP requests to the sentiment analysis service to get the sentiment score for each comment.

#### Sentiment Analysis Endpoint

The sentiment analysis service exposes an endpoint at `/sentiment` which accepts POST requests with the following JSON payload:

```json
{
  "text": "The comment text to analyze"
}
```

The service responds with a JSON array containing the sentiment label and score:

```json
[
  {
    "label": "POSITIVE",
    "score": 0.95
  }
]
```

### How Sentiment Analysis is Integrated

In the backend, the `chatbot_controller.go` file has been updated to include calls to the sentiment analysis service. When a customer submits a comment, the application sends the comment text to the sentiment analysis service and logs the result.

Here is the relevant code snippet from `chatbot_controller.go`:

```go
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
```

By incorporating sentiment analysis, we aim to provide a more comprehensive understanding of customer feedback, enabling more informed decisions to enhance customer satisfaction and product quality.

## Scope

The scope is limited to what was requested in the document, without the possibility of reviewing requirements and refining tasks. The limited execution time of 4 hours ends up leaving aside some implementations such as: unit testing, interface testing, integration testing, monitoring, API documentation, among other things. If it is necessary to implement, give me a few more hours.

## Contribution

Feel free to contribute to the project. Fork the repository, create a branch, and submit a pull request with your changes.
