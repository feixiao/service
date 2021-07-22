## JuiceFS快速入门

#### 简介
创建JuiceFS需要三步:
+ Redis保存元数据
+ 对象存储使用minio
+ JuiceFS Client

#### 安装
##### Redis
```shell
#  打开 --appendonly yes
sudo docker run -d --name redis_with_data \
	-v ~/volumes/redis/data:/data \
	-p 16379:16379 \
	--restart unless-stopped \
	redis:4.0.14 redis-server --appendonly yes
```
##### Object Storage
```shell
sudo docker run -d --name minio \
	-v  ~/volumes/minio-data:/data \
	-p 9000:9000 \
	--restart unless-stopped \
	minio/minio server /data
```

#### 安装JuiceFS客户端
```shell
wget https://github.com/juicedata/juicefs/releases/download/v0.15.2/juicefs-0.15.2-linux-amd64.tar.gz

sudo cp -rf ./juicefs /usr/local/bin

```


#### 创建JuiceFS文件系统
```shell
juicefs format \
	--storage minio \
	--bucket http://127.0.0.1:9000/pics \
	--access-key minioadmin \
	--secret-key minioadmin \
	redis://127.0.0.1:16379/1 \
	pics
```


#### 参考资料
+ [juicefs](https://github.com/juicedata/juicefs)
+ [《JuiceFS Quick Start Guide》](https://github.com/juicedata/juicefs/blob/main/docs/en/quick_start_guide.md)



