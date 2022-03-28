# syntax=docker/dockerfile:1

## Building Layer
FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o /hello-world


## Running Layer
FROM scratch
WORKDIR /

COPY --from=builder /hello-world .

EXPOSE 8090

USER nonroot:nonroot

ENTRYPOINT ["/hello-world"]
