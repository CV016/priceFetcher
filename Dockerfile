# Use the latest Golang version (1.21) for the base image
FROM golang:1.21-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /priceFetcher

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["/priceFetcher"]
