FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./go /go/src/github.com/user/monkey-ops/app
WORKDIR /go/src/github.com/user/monkey-ops/app

RUN go get ./
RUN go build

CMD app;
	
EXPOSE 8080
