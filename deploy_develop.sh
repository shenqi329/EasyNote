#!/bin/sh
docker ps -a |grep easynote |grep -v grep
if [ $? -ne 0 ]
then
docker run -it -p 80:80 -v /go/src/easynote:/go/src/easynote -v /go/src/easynote_data:/go/src/easynote_data --name easynote shenqi329/easynote /bin/bash
else
docker start -i easynote
fi



#docker run -it -p 80:80 -v /go/src:/go/src --name easynote golang /bin/bash

docker run -it -v /go/src:/go/src --name connect2 golang /bin/bash

docker run -it -v /go/src:/go/src --name connect3 golang /bin/bash

docker run -it -v /go/src:/go/src --name connect5 connect /bin/bash