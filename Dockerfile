FROM golang:1.22.8-bookworm

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod . 
COPY go.sum .

# Ensure the dependencies are downloaded
RUN go mod tidy 

# Copy the source code into the container
COPY . .

# Set build arguments
ARG MONGODB_URI
ARG JWT_SECRET
ARG DB_NAME

# Create .env file and add environment variables
RUN touch .env && \
    echo "MONGODB_URI=$MONGODB_URI" >> .env && \
    echo "JWT_SECRET=$JWT_SECRET" >> .env && \
    echo "DB_NAME=$DB_NAME" >> .env

# Build the Go application
RUN go build -o app .

# Expose the port
EXPOSE 8080

# Set the command to run the application
CMD ["./app"]
