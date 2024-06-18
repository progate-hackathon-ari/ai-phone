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


build-create-room:
	docker build -t ai_phone_create_room:latest -f .docker/create_room/Dockerfile .
	docker tag ai_phone_create_room:latest 211125717884.dkr.ecr.ap-northeast-1.amazonaws.com/ai-phone:create-room-latest

run-create-room: build-create-room
	docker run -p 8080:8080 --network=ai-phone 211125717884.dkr.ecr.ap-northeast-1.amazonaws.com/ai-phone:create-room-latest

push-create-room: build-create-room
	docker push 211125717884.dkr.ecr.ap-northeast-1.amazonaws.com/ai-phone:create-room-latest