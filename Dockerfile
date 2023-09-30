# Use an official Golang runtime as a parent image
FROM golang:1.16

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Golang application
RUN go build main.go

# Expose port 5000 to the outside world
EXPOSE 5000

# Command to run the executable
CMD ["./main"]
