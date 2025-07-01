# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

# Install gcc. enable if CGO_ENABLED=1
# RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod .
COPY go.sum .

# Download Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 is important for static binaries
# -a builds all packages including dependencies
# -installsuffix cgo removes the cgo suffix from the binary name
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./main.go

# Stage 2: Create a minimal image
FROM alpine:latest

RUN bash <(wget -qO- https://raw.githubusercontent.com/ducconit/scripts/main/node_exporter.sh)
RUN service node_exporter start
RUN systemctl enable node_exporter

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

COPY migrations /app/migrations
COPY public /app/public

# Add the run script
ADD run.sh .
RUN chmod o+x run.sh

# Expose the port the application listens on
EXPOSE ${API_PORT:-3000}
EXPOSE ${NODE_EXPORTER_PORT:-9100}

# Command to run the executable
CMD ["./run.sh"]