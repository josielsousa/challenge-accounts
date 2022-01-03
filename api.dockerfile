FROM gcr.io/distroless/static-debian10

WORKDIR /app

COPY challange-accounts /

COPY app/gateway/db/postgres/migrations /migrations

EXPOSE 3000

ENTRYPOINT [ "/challange-accounts" ]