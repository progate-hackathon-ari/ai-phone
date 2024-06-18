FROM golang:1.22.2 AS build
WORKDIR /build

COPY go.mod go.sum /

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main cmd/create_room/main.go

FROM public.ecr.aws/lambda/go:latest AS runner

COPY --from=build /build/main main

CMD ["main"]