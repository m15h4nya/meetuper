FROM golang:alpine3.18 as build

COPY . /meetupper
RUN cd /meetupper && go build -o service main.go

FROM debian:latest as run

RUN apt-get update && apt-get install -y ca-certificates openssl
ARG cert_location=/usr/local/share/ca-certificates
RUN update-ca-certificates

RUN mkdir -p /opt/meetupper
COPY --from=build /meetupper/service /opt/meetupper/service
COPY --from=build /meetupper/config/config.toml /opt/meetupper/config/config.toml
ENV TZ=Europe/Moscow
WORKDIR /opt/meetupper
CMD ["./service"]