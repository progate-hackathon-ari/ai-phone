FROM golang:1.22.2 AS builder
WORKDIR /build

COPY go.mod go.sum /

RUN go mod download

COPY . .

RUN go build -o main cmd/app/main.go

FROM ubuntu:slim AS runner
WORKDIR /app

COPY --from=builder main main

CMD ["./main"]