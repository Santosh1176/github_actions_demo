FROM golang:1.19 AS build

WORKDIR /app

COPY . /app

RUN go build main

FROM alpine:latest

WORKDIR /root

COPY --from=build /pkg /root

ENTRYPOINT [ "./main"]