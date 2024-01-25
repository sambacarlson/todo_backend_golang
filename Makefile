ENV_VARS := $(shell sed 's/ /\\ /g' .env)
DB_URL := $(shell echo "${ENV_VARS}" | grep -o "DB_URL:[^ ]*" | cut -d ':' -f 2-)
DB_PORT := $(shell echo "${ENV_VARS}" | grep -o "DB_PORT:[^ ]*" | cut -d ':' -f 2-)
DB_PASSWORD := $(shell echo "${ENV_VARS}" | grep -o "DB_PASSWORD:[^ ]*" | cut -d ':' -f 2-)
DB_VOLUME_PATH := $(shell echo "${ENV_VARS}" | grep -o "DB_VOLUME_PATH:[^ ]*" | cut -d ':' -f 2-)
DB_IMAGE_NAME := $(shell echo "${ENV_VARS}" | grep -o "DB_IMAGE_NAME:[^ ]*" | cut -d ':' -f 2-)
DB_CONTAINER_NAME := $(shell echo "${ENV_VARS}" | grep -o "DB_CONTAINER_NAME:[^ ]*" | cut -d ':' -f 2-)
API_IMAGE_NAME := $(shell echo "${ENV_VARS}" | grep -o "API_IMAGE_NAME:[^ ]*" | cut -d ':' -f 2-)

export DATABASE_URL=$(DB_URL)


tidy:
	go mod tidy

lint:
	gofmt -s -w .
	golangci-lint run .

fix:
	golangci-lint run --fix

run: tidy
	go run ./main.go


.PHONY: build-image
build-image:
	docker build -t $(API_IMAGE_NAME) .

create-migration:
	migrate create -ext sql -dir db/migrations -digits 6 -seq $(name); 

run-db:
	docker container prune -f
	docker run --name $(DB_CONTAINER_NAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -p $(DB_PORT):$(DB_PORT) -v $(DB_VOLUME_PATH):/mytodo_data -it $(DB_IMAGE_NAME)

start-db:
	docker start $(DB_CONTAINER_NAME)

migrate-up:
	migrate -path ./db/migrations -database "$$DATABASE_URL" up
