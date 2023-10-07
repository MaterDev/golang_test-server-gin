# Run the server using air for hot reloading
dev:
	air

# Build the application
build:
	go build -o main .

# Run the application locally
run:
	./main

# Build the Docker image
docker-build:
	docker build -t golang_test-server-gin .

# Run the Docker container
docker-run: docker-build
	docker run -p 8080:8080 golang_test-server-gin