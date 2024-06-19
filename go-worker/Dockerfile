# Base image for building Go app
FROM golang:latest as builder

# Set working directory
WORKDIR /go/src/app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Base image for running the app
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the built binary from builder
COPY --from=builder /go/src/app/main .

# Expose port 5000
EXPOSE 5000

# Command to run the application
CMD ["./main"]
