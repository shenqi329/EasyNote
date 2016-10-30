# docker@default设置

##共享设置

###virtualbox 创建共享文件夹，命名为easynote

###创建本地目录
####sudo mkdir -p /go/src/easynote

###设置共享文件夹，shareName virtualbox设置共享的文件夹

####sudo mount -t vboxsf shareName /go/src/easynote
####如:sudo mount -t vboxsf easynote /go/src/easynote
