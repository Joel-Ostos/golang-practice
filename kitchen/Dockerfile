# Dockerfile para el microservicio de cocina
FROM golang:1.18-alpine

RUN apk update && apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY persistence/data /app/persistence/data

COPY . .

RUN go build -o main .

CMD ["./main"]

