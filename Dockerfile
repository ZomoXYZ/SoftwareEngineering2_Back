# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

WORKDIR /
COPY . .

# build-base is required for gcc
# bash is just for debugging
RUN apk add build-base bash

RUN go mod download
RUN go build -o main.o .

EXPOSE 8080
CMD ["/main.o"]
