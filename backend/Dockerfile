# Use the official Golang image to create a build artifact.
# This is the build stage.
FROM golang:1.18-alpine AS builder

# Install git and other dependencies
RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY wait-for-it.sh /root/wait-for-it.sh

# Ensure the wait-for-it script has execution permissions
RUN chmod +x /root/wait-for-it.sh

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable, waiting for the database and sentiment analysis service
CMD /root/wait-for-it.sh db 3306 -- /root/wait-for-it.sh sentiment-analysis 5000 -- ./main
