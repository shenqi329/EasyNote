FROM golang

MAINTAINER shenqi329 <shenqi329@163.com>

RUN apt-get update
RUN apt-get -y install mongodb

COPY .  /usr/share/note/

WORKDIR /usr/share/note/

RUN ls -al

RUN chomd 777 ./main

RUN  ./main

EXPOSE 1323


