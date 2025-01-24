FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN go build -o main .

FROM golang:1.23
WORKDIR /root/
COPY --from=builder /app/main .
COPY .env .
EXPOSE 8089
CMD ["./main"]
