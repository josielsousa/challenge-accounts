# Build image app from golang:alpine image base.
FROM golang:1.14.5-alpine AS build

WORKDIR /app 

COPY . /app

RUN apk add --no-cache --update build-base

RUN make build 

# Steps to run app from nginx imagem base.
FROM nginx:alpine

WORKDIR /app

COPY --from=build /app/bin/service bin/service

EXPOSE 3000

CMD ["./bin/service"]


