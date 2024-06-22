migrate-up:
	go run cmd/migrate/main.go -e .env -f cmd/migrate/schema up

migrate-down:
	go run cmd/migrate/main.go -e .env -f cmd/migrate/schema down

migrate-up-test:
	go run cmd/migrate/main.go -e .env.test -f cmd/migrate/schema up

migrate-down-test:
	go run cmd/migrate/main.go -e .env.test -f cmd/migrate/schema down

gen-model: migrate-up
	go run gorm.io/gen/tools/gentool@latest -c ./gen.tool.yaml


build-app:
	docker build -t ai-phone -f .docker/app/Dockerfile .
	docker tag ai-phone:latest 830373977162.dkr.ecr.us-east-1.amazonaws.com/ai-phone:latest

run-create-room: build-create-room
	docker run -p 8080:8080 --network=ai-phone 211125717884.dkr.ecr.ap-northeast-1.amazonaws.com/ai-phone:create-room-latest

push-app: build-app
	docker push 830373977162.dkr.ecr.us-east-1.amazonaws.com/ai-phone:latest