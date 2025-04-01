FROM golang:1.21-alpine

WORKDIR /app

# Copy go mod first
COPY go.mod ./

# Initialize module and download dependencies
RUN go mod tidy

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Expose port 8083
EXPOSE 8083

# Copy static files and templates
COPY static/ ./static/
COPY templates/ ./templates/

# Command to run the application
CMD ["./main"]
