FROM golang:1.21.3

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o ./build/main ./cmd/main.go
