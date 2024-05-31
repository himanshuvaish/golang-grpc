# Use the official Golang image to create a build artifact
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Ensure all dependencies are included in go.sum
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o server ./server
RUN go build -o client ./client

# Start a new stage from scratch
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary files from the previous stage
COPY --from=builder /app/server /root/
COPY --from=builder /app/client /root/

# Expose port 80 to the outside world
EXPOSE 80

# Entry point allows passing arguments to decide whether to run server or client
ENTRYPOINT ["sh", "-c"]
CMD ["./server"]