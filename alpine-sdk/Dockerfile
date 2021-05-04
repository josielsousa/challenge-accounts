FROM golang:1.14.5-alpine AS build

WORKDIR /app 

COPY app /app

RUN apk add --no-cache --update make ca-certificates openssl alpine-sdk 

RUN make build 


FROM golang:1.14.5-alpine

WORKDIR /app

COPY --from=build /app/bin/service bin/service

EXPOSE 3000

CMD ["./bin/service"]


