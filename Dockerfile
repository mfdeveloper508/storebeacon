FROM golang:1.9-alpine

ARG app_env
ENV APP_ENV $app_env

RUN apk add --update git curl

ADD . /go/src/github.com/storebeacon/backend
WORKDIR /go/src/github.com/storebeacon/backend

RUN go build -o fupp-api main.go

COPY bin/docker-entrypoint.sh /usr/local/bin
RUN ["chmod", "+x", "/usr/local/bin/docker-entrypoint.sh"]

COPY bin/run-tests.sh /usr/local/bin
RUN ["chmod", "+x", "/usr/local/bin/run-tests.sh"]

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]

EXPOSE 8080
