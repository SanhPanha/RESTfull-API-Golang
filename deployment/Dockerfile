# Stage 1: Build the Go application
FROM golang:1.23-alpine AS build

WORKDIR /app

# Install gcc and musl-dev to support cgo
RUN apk add --no-cache gcc musl-dev

# Copy go.mod and go.sum files first to take advantage of Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory to the container
COPY . .

# Build the Go application with cgo enabled
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main ./cmd/app/main.go

# Stage 2: Run the application
FROM alpine:latest

# Install SQLite3 in the final image
RUN apk add --no-cache sqlite-libs

WORKDIR /app

# Create the db directory in the container
RUN mkdir -p /app/db

# Copy the compiled binary from the build stage
COPY --from=build /app/main .

# Copy the SQLite database file from the build context
COPY db/test.db /app/db/test.db

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the Go app
CMD ["./main"]