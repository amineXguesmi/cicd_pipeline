FROM golang:1.22.3-bookworm

WORKDIR /app

COPY go.mod . 

COPY go.sum .

RUN go mod tidy

COPY . .

ARG MONGODB_URI
ARG JWT_SECRET
ARG DB_NAME

RUN touch .env && \
    echo "MONGODB_URI=$MONGODB_URI" >> .env && \
    echo "JWT_SECRET=$JWT_SECRET" >> .env && \
    echo "DB_NAME=$DB_NAME" >> .env

RUN go build -o app .

EXPOSE 8080

CMD ["go" , "run" , "."]
