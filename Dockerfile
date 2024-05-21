# Start from the official Golang image to build the binary.
FROM golang:1.21.2 as builder

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum to download dependencies.
COPY go.mod go.sum ./

# Download dependencies.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myApp .

# Use a small Alpine Linux image to run the application.
FROM alpine:latest

# Install certificates for HTTPS.
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container.
WORKDIR /root/

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/myApp .
COPY --from=builder /app/config.toml .


# Run the binary with ENTRYPOINT.
ENTRYPOINT ["./myApp"]