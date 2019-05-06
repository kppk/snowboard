FROM golang:1.12-stretch AS builder
MAINTAINER Alif Rachmawadi <subosito@bukalapak.com>

WORKDIR /app
COPY . .
RUN make build

FROM debian:stretch-slim
COPY --from=builder /app/snowboard /usr/local/bin

RUN apt-get -y update \
 && apt-get -y install --no-install-recommends inotify-tools \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /doc
VOLUME /doc
EXPOSE 8088 8087

ENTRYPOINT ["snowboard"]
