FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Install the required packages
RUN apk update && apk add --no-cache git
RUN go mod download

# Build the application
RUN go build -o main .

# Expose the port
EXPOSE 8000

# Run the application
CMD ["./main"]
