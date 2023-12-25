FROM golang:1.21-alpine as base

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

COPY ./go.mod ./go.sum ${ROOT}/

RUN go mod download

CMD [ "/bin/sh" ]