test:
	go test ./...

get-u:
	go get -u ./...

build:
	docker compose build

run: build
	docker compose up

rund: build
	docker compose up -d

migrate-up:
	go run cmd/migrate/main.go -e .env -f cmd/migrate/schema up

migrate-down:
	go run cmd/migrate/main.go -e .env -f cmd/migrate/schema down

migrate-up-test:
	go run cmd/migrate/main.go -e .env.test -f cmd/migrate/schema up

migrate-down-test:
	go run cmd/migrate/main.go -e .env.test -f cmd/migrate/schema down

dsn=postgres://postgres:postgres@localhost:5432/hal_cinema

gen-model: migrate-up
	go run gorm.io/gen/tools/gentool@latest -c ./gen.tool.yaml

seeder:
	go run cmd/seeder/main.go -e .env -f cmd/seeder/seeder.json