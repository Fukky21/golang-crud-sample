FROM golang:latest

ENV APP_ROOT /go/src/app

RUN mkdir -p $APP_ROOT

WORKDIR $APP_ROOT
