
FROM golang:1.23-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download


COPY . .


RUN go build -o hello-api .


FROM alpine:latest  

WORKDIR /app


COPY --from=builder /app/hello-api .

EXPOSE 8084


CMD ["./hello-api"]
