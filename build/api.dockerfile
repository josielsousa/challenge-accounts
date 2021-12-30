FROM gcr.io/distroless/static-debian10

WORKDIR /app

COPY challange-accounts /app/challange-accounts

EXPOSE 3000

ENTRYPOINT [ "/app/challange-accounts" ]