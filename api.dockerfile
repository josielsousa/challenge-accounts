FROM gcr.io/distroless/static-debian10

ADD --chown=nonroot:nonroot build/challange-accounts /challange-accounts

COPY --chown=nonroot:nonroot app/gateway/db/postgres/migrations /migrations

EXPOSE 3000

ENTRYPOINT [ "/challange-accounts" ]