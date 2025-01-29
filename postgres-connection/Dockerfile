# Use an official Go image
FROM golang:1.23

# Set the working directory
WORKDIR /

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
