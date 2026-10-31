FROM golang:1.27.1-alpine3.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./business/ ./business
RUN go build -o business_service ./business/cmd

FROM alpine:3.24
WORKDIR /app
COPY --from=builder /app/business_service .
RUN chmod +x ./business_service
EXPOSE 8888
CMD ["./business_service"]