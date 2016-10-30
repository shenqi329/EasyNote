FROM golang

MAINTAINER shenqi329 <shenqi329@163.com>

RUN apt-get update
RUN apt-get -y install mongodb

COPY .  /go/src/easynote/

WORKDIR /go/src/easynote/

RUN go get github.com/labstack/echo
RUN go get gopkg.in/mgo.v2
RUN go get github.com/dgrijalva/jwt-go
RUN go build main.go

EXPOSE 1323

CMD ./main


