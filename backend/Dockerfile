FROM debian:12.1-slim

RUN apt update 
RUN apt show golang
RUN apt install golang ca-certificates -y
RUN apt install host -y

WORKDIR /app

COPY . .

RUN go mod download

RUN go build cmd/main.go

RUN chmod 777 main

CMD ["./main"]