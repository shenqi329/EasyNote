FROM golang

MAINTAINER shenqi329 <shenqi329@163.com>

RUN apt-get update
RUN apt-get -y install mongodb supervisor

COPY .  /go/src/easynote/
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

WORKDIR /go/src/easynote/

RUN go get github.com/labstack/echo
RUN go get gopkg.in/mgo.v2
RUN go get github.com/dgrijalva/jwt-go
RUN go build main.go

EXPOSE 80

CMD ["/usr/bin/supervisord"]


