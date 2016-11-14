#!/bin/sh
docker ps -a |grep easynote |grep -v grep
if [ $? -ne 0 ]
then
docker run -it -p 80:80 -v /go/src/easynote:/go/src/easynote -v /go/src/easynote_data:/go/src/easynote_data --name easynote shenqi329/easynote /bin/bash
else
docker start -i easynote
fi



#docker run -it -p 80:80 -p 27017:27017 -p 8081:8081 -p 8082:8082 -p 8888:22 -v /go/src/easynote:/go/src/easynote -v /go/src/easynote_data:/go/src/easynote_data --name easynote shenqi329/easynote /bin/bash
