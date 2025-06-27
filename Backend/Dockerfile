# Use Go Alpine
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

# COPY .env .env

# Copy all source files
COPY . .

# Build the Go app
RUN go build -o main ./cmd/server/main.go

# Run the binary
CMD ["./main"]
