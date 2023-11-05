# golang_test-server-gin

Golang test server (with Gin)

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
