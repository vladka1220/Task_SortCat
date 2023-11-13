
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./


RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

COPY --from=builder /app/main /main


CMD ["./main"]
