FROM golang:1.21.4 AS builder

WORKDIR /app

COPY go.mod go.sum .env ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest  

ENV MONGODB_URI="mongodb://localhost:27017"

WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]