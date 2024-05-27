ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# Specify the directory where main.go is located
WORKDIR /usr/src/app/cmd/server

# Build the Go application
RUN go build -v -o /run-app .

FROM debian:bookworm

# Copy the compiled binary to the final image
COPY --from=builder /run-app /usr/local/bin/

# Copy the web directory to the final image
COPY --from=builder /usr/src/app/web /usr/src/app/web

# Set the working directory to the location of the web directory
WORKDIR /usr/src/app

# Expose ports
EXPOSE 80
EXPOSE 443

CMD ["run-app"]
