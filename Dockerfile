FROM golang:1.19 AS build

WORKDIR /src

COPY go.mod go.sum /src/

RUN go mod download

COPY docker /src/

ARG COMMIT_HASH

RUN CGO_ENABLED=0 GOOS=linux go build -a \
    -ldflags="-w -s" -installsuffix cgo -o /app/send_email cmd/send_email/main.go

FROM alpine:3.16 AS release

WORKDIR /app

COPY --from=build /app/ /app/
