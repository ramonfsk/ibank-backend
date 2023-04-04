# syntax=docker/dockerfile:1

# Build
FROM golang:1.16-alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
# RUN CGO_ENABLED=0 go test -v

RUN go build -o ibank-server main.go

# Deploy
FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/ibank-server /app/ibank-server

EXPOSE 8010

CMD [ "./ibank-server" ]