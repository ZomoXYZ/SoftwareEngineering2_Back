# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

WORKDIR /
COPY . .

# don't copy any database files
RUN rm -f /*.db

# required for gcc
RUN apk add build-base

RUN go mod download
RUN go build -o main.o .

EXPOSE 8080
CMD ["/main.o"]
