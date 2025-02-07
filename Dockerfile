FROM golang:1.23 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main . && ls -lah main

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main .

RUN chmod +x main && ls -lah main

EXPOSE 8080

CMD ["./main"]
