FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go mod download

RUN go build -o main payment-manager/cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

COPY .env .

EXPOSE 8081

CMD ["./main"]
