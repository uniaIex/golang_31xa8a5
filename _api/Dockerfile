FROM golang:latest

WORKDIR /golang_31xa8o3

# Download Go modules
COPY . .
RUN go mod download

# Build
RUN go build -o ./golang_31xa8o3 .

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8002

ENV GIN_MODE=release
# Run
ENTRYPOINT ["./golang_31xa8o3"]