# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Aleksandar Pavlovic"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ./food-service/go.mod ./food-service/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy everything from the current directory to the Working Directory inside the container
COPY ./food-service/ .

# Try running go build command separately to troubleshoot
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]
