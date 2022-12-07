FROM golang:1.19.3-alpine3.17 as base

FROM base as dev

RUN apk update && apk upgrade && apk add curl

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api
CMD ["air"]