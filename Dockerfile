FROM golang:1.21-alpine as base

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

COPY ./src/app/go.mod ./src/app/go.sum ${ROOT}/

COPY . ${ROOT}/

RUN go mod download

CMD [ "/bin/sh" ]
