FROM golang:1.22.2 AS builder
WORKDIR /build

COPY go.mod go.sum /

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/app/main.go

FROM gcr.io/distroless/base-debian11 AS runner
WORKDIR /app

COPY --from=builder build/main main

CMD ["./main"]