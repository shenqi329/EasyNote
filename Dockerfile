FROM golang

MAINTAINER shenqi329 <shenqi329@163.com>

RUN apt-get update
RUN apt-get -y install mongodb

COPY .  /usr/share/note/

WORKDIR /usr/share/note/

RUN go get github.com/labstack/echo
RUN go get gopkg.in/mgo.v2
RUN go get github.com/dgrijalva/jwt-go
RUN go build main.go
RUN  ./main

EXPOSE 1323


