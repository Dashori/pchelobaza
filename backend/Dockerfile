FROM golang:1.21.3 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./build/main ./cmd/main.go

# FROM alpine:3.18

# WORKDIR /app

# COPY --from=build /build/main .

# CMD ["./main"]