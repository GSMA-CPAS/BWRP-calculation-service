FROM golang:1.23-alpine3.20

COPY . /app

WORKDIR /app

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -d /app/cmd/calculation-api/

RUN go build -o main /app/cmd/calculation-api/

EXPOSE 8080

CMD ["./main"]
