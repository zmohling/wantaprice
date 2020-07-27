FROM golang:1.14-alpine as builder

RUN apk add --no-cache git

COPY . /app

WORKDIR /app

RUN GOARCH=amd64 \
    GOOS=linux \
    go build -v -o identityservice

# ---------------------------------------------------------------------
#  The second stage container, for running the application
# ---------------------------------------------------------------------
FROM alpine:latest
COPY --from=builder /app/identityservice /identityservice

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["/identityservice"]