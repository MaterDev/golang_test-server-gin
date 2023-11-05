# Golang test server (with Gin)

## Project Overview

This project is a spike to understand the basics of creating a Go/Gin server. It includes basic API endpoints, with a focus on a simple health check, and is designed to illustrate basic server architecture, containerization, and tooling for development.

## Prerequisites

- Go version `1.21.x`
- Docker (for containerization)
- Make (for running Makefile commands)
- Air (for live reloading)

## Setup and Installation

1. Clone the repository:

   ```sh
   git clone git@github.com:MaterDev/golang_test-server-gin.git
   ```

2. Navigate to the project directory:

   ```sh
   cd golang_test-server-gin
   ```

3. Install dependencies:

   ```sh
   go mod tidy
   ```

## Running the Server Locally

To run the server locally, execute:

```sh
make run
```

For live reloading, use:

```sh
make dev
```

> *This requires the Air to be installed. If you don't have it, you can follow the setup instructions [here](https://github.com/cosmtrek/air).*

Alternativiely you can build and then start the server with:

```sh
go build -o build/main src/cmd/server/main.go
./build/main
```

## Using the Makefile

The Makefile includes commands for building, running, and cleaning up the project:

- `make build`: Compiles the application and generates a binary.
- `make run`: Runs the server directly.
- `make clean`: Removes compiled binaries and other artifacts.

## Containerization with Docker

Build the Docker image using:

```sh
docker build -t go-gin-server .
```

Run the server in a Docker container:

```sh
docker run -p 8080:8080 go-gin-server
```

You can also use the Makefile to build and run the Docker image:

- `make docker-build`: Builds the Docker image.
- `make docker-run`: Runs the server in a Docker container.

## Endpoints

- `GET /healthcheck`: Returns the health status of the server.

## Directory Structure

```bash
golang-gin-server/
├── Dockerfile             # Instructions for Docker to build the application's container image
├── Makefile               # Defines set of tasks to be executed. Commonly used to compile/build the application.
├── README.md              # Documentation about this project, how to build, run, test, etc.
├── build                  # Folder containing build artifacts, often .gitignored
│   └── main               # Compiled binary after building the Go application
├── go.mod                 # Go module definition with module name and dependency requirements
├── go.sum                 # Contains the expected cryptographic checksums of the content of specific versions of dependencies
├── src                    # Source directory where the primary application code is located
│   ├── api                # Houses the API specific components such as handlers, middleware, and route definitions
│   │   ├── handlers       # Contains handler functions that correspond to API endpoints
│   │   │   └── healthcheck.handler.go # Handler for healthcheck endpoint, returning the server status
│   │   ├── middlewares    # Contains middleware components, which are executed before/after handlers
│   │   │   ├── errorhandler.go        # Middleware for error handling, could capture and format errors
│   │   │   └── logger.go              # Middleware for logging, handles log setup and log entries
│   │   └── routes         # Contains route definitions, connecting URLs to handlers
│   │       └── healthcheck.route.go   # Route setup for the healthcheck endpoint
│   └── cmd                # Contains the entry point of the application
│       └── server         # Entry point for the server application specifically
│           └── main.go    # The main application file, where the server is initialized and run
└── tmp                    # Temporary files, commonly used for logs or other temporary data
    ├── air.log            # Log file for the live reloading tool or another process
    └── main               # Sometimes a temporary binary is stored here during development
```
