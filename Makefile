IMAGE_NAME := mytodo_api
CONTAINER_NAME := mytodo_api_container
IMAGE_NAME := myapi
CONTAINER_NAME := myapi-container

tidy:
	go mod tidy

lint:
	gofmt -s -w .
	golangci-lint run .

fix:
	golangci-lint run --fix

run: tidy
	go run ./main.go

.PHONY: clean
clean:
	-docker rm -f $(CONTAINER_NAME)

.PHONY: build-image
build-image:
	docker build -t $(IMAGE_NAME) .

.PHONY: run-container
run-container:
	-docker rm -f $(CONTAINER_NAME)
	docker run -p 8080:8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)
