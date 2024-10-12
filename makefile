# Define variables
IMAGE_NAME := my-go-app
CONTAINER_NAME := my-go-app-container
PORT := 3000

# Targets
.PHONY: build run stop clean

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME) .

# Run the Docker container
run: build
	docker run -d --name $(CONTAINER_NAME) -p $(PORT):$(PORT) $(IMAGE_NAME)

# Stop the running container
stop:
	docker stop $(CONTAINER_NAME)

# Remove the container and the image
clean: stop
	docker rm $(CONTAINER_NAME)
	docker rmi $(IMAGE_NAME)

# Help
help:
	@echo "Makefile commands:"
	@echo "  make build   - Build the Docker image"
	@echo "  make run     - Run the Docker container"
	@echo "  make stop    - Stop the Docker container"
	@echo "  make clean   - Stop and remove the container and image"
