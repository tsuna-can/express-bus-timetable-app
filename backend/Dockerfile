FROM golang:1.23.2-bullseye as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and download dependencies
COPY ./backend/go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

FROM gcr.io/distroless/base-debian12

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go application from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]