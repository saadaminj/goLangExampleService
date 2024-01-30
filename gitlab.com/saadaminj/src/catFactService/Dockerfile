# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go service inside the container
RUN go build -o mygoservice .

# Run the Go service when the container starts
CMD ["./mygoservice"]
