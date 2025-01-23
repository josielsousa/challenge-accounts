FROM golang:1.23.1-alpine AS builder

# Load the public keys from github and configure ssh url.
RUN apk update && apk add openssh git tzdata ca-certificates && update-ca-certificates
RUN mkdir -p -m 0700 ~/.ssh && ssh-keyscan github.com > ~/.ssh/known_hosts
RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

RUN adduser -D -H -h "/nonexistent" -s "/sbin/nologin" -g "" -u "10001" "appuser"

# Copy files and fetch dependencies.
WORKDIR $GOPATH/src/challange-accounts/
ENV GOPRIVATE=github.com/josielsousa/*

COPY go.mod .
COPY go.sum .

RUN --mount=type=ssh go mod download
COPY app ./app
COPY cmd ./cmd
COPY types ./types

# Build the binary.
ARG GIT_BUILD_TIME
ARG GIT_COMMIT
ARG GIT_TAG
RUN GOOS=linux GOARCH=amd64 go build -a -o /go/bin/challange-accounts \
    -ldflags="-w -s -X main.BuildTime=$GIT_BUILD_TIME -X main.BuildCommit=$GIT_COMMIT -X main.BuildTag=$GIT_TAG" \
    ./cmd

# Create a minimal image.
FROM gcr.io/distroless/static-debian10

# Import from builder.
COPY --chown=nonroot:nonroot --from=builder /etc/group /etc/group
COPY --chown=nonroot:nonroot --from=builder /etc/passwd /etc/passwd
COPY --chown=nonroot:nonroot --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=nonroot:nonroot --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the binary.
COPY --from=builder /go/bin/challange-accounts /go/bin/challange-accounts

# Use the unprivileged user.
USER appuser:appuser

EXPOSE 3000

# Run the binary.
ENTRYPOINT ["/go/bin/challange-accounts"]
