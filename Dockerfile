# # Start from the official Golang image to build the binary.
# FROM golang:1.21.2 as builder

# # Set the working directory inside the container.
# WORKDIR /app

# # Copy go.mod and go.sum to download dependencies.
# COPY go.mod go.sum ./

# # Download dependencies.
# RUN go mod download

# # Copy the source code into the container.
# COPY . .

# # Build the application.
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myApp .

# # Use a small Alpine Linux image to run the application.
# FROM alpine:latest

# # Install certificates for HTTPS.
# RUN apk --no-cache add ca-certificates

# # Set the working directory inside the container.
# WORKDIR /root/

# # Copy the pre-built binary file from the previous stage.
# COPY --from=builder /app/myApp .
# # COPY --from=builder /app/config.toml .


# # Run the binary with ENTRYPOINT.
# CMD ["./myApp"]

FROM golang:1.21.2 as builder
# RUN apt-get update && apt-get -y upgrade && apt-get install -y upx
COPY . /build/app
WORKDIR /build/app

RUN go get ./... && go build -ldflags "-s -w" -trimpath -o backend main.go
# RUN upx backend && upx -t backend

# 2nd stage, create a user to copy, and install libraries needed if connecting to upstream TLS server
# we don't want the /lib and /lib64 from the go container cause it has more than we need.
FROM debian:11 AS ssl
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && apt-get -y upgrade && apt-get install -y ca-certificates && \
    addgroup --gid 26657 --system backend && adduser -uid 26657 --ingroup backend --system --home /var/lib/backend backend

# 3rd and final stage, copy the minimum parts into a scratch container, is a smaller and more secure build. This pulls
# in SSL libraries and CAs so Go can connect to TLS servers.
FROM scratch
COPY --from=ssl /etc/ca-certificates /etc/ca-certificates
COPY --from=ssl /etc/ssl /etc/ssl
COPY --from=ssl /usr/share/ca-certificates /usr/share/ca-certificates
COPY --from=ssl /usr/lib /usr/lib
COPY --from=ssl /lib /lib
COPY --from=ssl /lib64 /lib64

COPY --from=ssl /etc/passwd /etc/passwd
COPY --from=ssl /etc/group /etc/group
COPY --from=ssl --chown=backend:backend /var/lib/backend /var/lib/backend

COPY --from=builder /build/app/backend /bin/backend
COPY --from=builder /build/app/example-config.toml /var/lib/backend

USER backend
WORKDIR /var/lib/backend

ENTRYPOINT ["/bin/backend"]