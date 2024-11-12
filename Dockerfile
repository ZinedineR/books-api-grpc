# syntax=docker/dockerfile:1

#FROM golang:1.22 alpine
FROM golang:1.22-alpine as build

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o books-api-user ./cmd/web/user
RUN go build -o books-api-author ./cmd/web/author
RUN go build -o books-api-category ./cmd/web/category
RUN go build -o books-api-book ./cmd/web/book

FROM alpine:edge

WORKDIR /app
COPY --from=build /app/books-api-user .
COPY --from=build /app/books-api-author .
COPY --from=build /app/books-api-category .
COPY --from=build /app/books-api-book .
