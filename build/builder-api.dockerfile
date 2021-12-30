# Build image app from golang:alpine image base.
FROM golang:1.17-alpine AS builder

WORKDIR /app 

COPY . /app

RUN apk add --no-cache ca-certificates build-base

RUN make compile  

# Steps to run app from distroless imagem base.
FROM gcr.io/distroless/static-debian10

WORKDIR /app

COPY challange-accounts /app/challange-accounts

EXPOSE 3000

ENTRYPOINT [ "/app/challange-accounts" ]
