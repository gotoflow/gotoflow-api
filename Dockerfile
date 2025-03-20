# Use the official Golang image as the base image
FROM golang:1.23

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Command to run the application with Air
EXPOSE 3333
CMD ["air", "-c", ".air.toml"]