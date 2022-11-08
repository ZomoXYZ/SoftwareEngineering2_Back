# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

WORKDIR /
COPY . .

# don't copy any database files
RUN rm -f /*.db

# build-base is required for gcc
# bash is just for debugging
RUN apk add build-base bash

RUN go mod download
RUN go build -o main.o .

EXPOSE 8080
CMD ["/main.o"]
