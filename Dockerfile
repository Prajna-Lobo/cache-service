FROM golang:1.15-alpine

ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR  $GOPATH/src/cache-service

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy everything from the current directory to the PWD inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# This container exposes port 8080 to the outside world
EXPOSE 8081

CMD ["./main"]