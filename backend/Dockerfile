FROM golang:1.21.3

WORKDIR /app

COPY . .

RUN go mod download && \
    go build -o ./build/main ./cmd/main.go
