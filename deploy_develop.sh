#!/bin/sh
docker run -it -p 80:80 -v /go/src/easynote:/go/src/easynote -v /go/src/easynote_data:/go/src/easynote_data --name easynote shenqi329/easynote /bin/bash