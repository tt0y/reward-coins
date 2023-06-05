# Stage 1: Build the Go application
FROM golang:1.20 AS go_builder

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the application source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Expose the necessary ports
EXPOSE 8080

# Start the Go application
CMD ["./app"]
