FROM golang:1.17-alpine as build

# Set the Current Working Directory inside the container
WORKDIR /app/safesky-ws

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/safesky-ws .

# This container exposes port 8080 to the outside world
EXPOSE 8000

# Run the binary program produced by `go install`
CMD ["./out/safesky-ws"]

FROM alpine:3.14.2 as production

WORKDIR /app/safesky-ws

COPY --from=base /app/safesky-ws/out/safesky-ws /bin/safesky-ws