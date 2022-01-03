# Build image app from golang:alpine image base.
FROM golang:1.17-alpine AS builder

WORKDIR /app 

COPY . /app

RUN apk add --no-cache ca-certificates build-base

RUN make compile  

# Steps to run app from distroless imagem base.
FROM gcr.io/distroless/static-debian10

WORKDIR /app

COPY --from=builder /build/challange-accounts /

COPY app/gateway/db/postgres/migrations /migrations

EXPOSE 3000

ENTRYPOINT [ "/challange-accounts" ]
