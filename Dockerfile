FROM golang:1.15-alpine3.13

COPY . /app

WORKDIR /app

RUN go mod download

RUN go get -u github.com/swaggo/swag/cmd/swag

RUN swag init -d /app/cmd/calculation-api/

RUN go build -o main /app/cmd/calculation-api/

EXPOSE 8080

CMD ["./main"]
