FROM golang:1.27.1-alpine3.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./review/ ./review
RUN go build -o review_service ./review/cmd

FROM alpine:3.24
WORKDIR /app
COPY --from=builder /app/review_service .
RUN chmod +x ./review_service
EXPOSE 8888
CMD ["./review_service"]