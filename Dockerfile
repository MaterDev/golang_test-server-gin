# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:1.20 as builder

# Copy local code to the container image.
WORKDIR /app
COPY . .

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main

# Use a minimal image to run the binary.
FROM gcr.io/distroless/base-debian10
COPY --from=builder /go/bin/main /go/bin/main

# Run the web service on container startup.
ENTRYPOINT ["/go/bin/main"]