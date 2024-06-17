FROM golang:1.22.2 AS migrate
WORKDIR /migrate

CMD ["go","run","cmd/migrate/main.go","-f","cmd/migrate/schema"]