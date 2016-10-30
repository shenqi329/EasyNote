#!/bin/sh
docker run -it -p 80:80 -v /go/src/easynote:/go/src/easynote --name easynote shenqi329/easynote /bin/bash