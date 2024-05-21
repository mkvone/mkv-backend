FROM golang:1.21.2 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myApp .

FROM alpine:latest

# Install certificates for HTTPS.
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container.
WORKDIR /root/
COPY --from=builder /app/myApp .
COPY --from=builder /app/config.yml .


ENTRYPOINT [""]