# docker@default设置

##共享设置

###virtualbox 创建共享文件夹，命名为easynote，用来存在代码，代码修改后在容器中也同步修改，便于编译、调试

###virtualbox 创建共享文件夹，命名为easynote_data，用来存在数据

###创建本地目录
####sudo mkdir -p /go/src/easynote
####sudo mkdir -p /go/src/easynote_data

###设置共享文件夹，shareName virtualbox设置共享的文件夹

####sudo mount -t vboxsf shareName /go/src/easynote
####如:sudo mount -t vboxsf easynote /go/src/easynote
####如:sudo mount -t vboxsf easynote_data /go/src/easynote_data

mount -t vboxsf easynote_data /go/src/easynote_data -o rw,uid=1000,gid=50,mode=0755

##开启容器
###进入/go/src/easynote文件夹，执行 sh deploy_develop.sh，此时进入了容器的控制台界面
###如果提示 easynote 容器已经存在，执行 docker rm -f easynote删除容器

###开启mongodb
#### nohup  mongod --config ./mongodb.conf &

