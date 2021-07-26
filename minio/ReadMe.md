## Minio 快速入门


#### 功能

#### 安装
```shell
# 稳定版本
docker pull minio/minio

docker run -p 9000:9000 minio/minio server /data


docker run -p 9000:9000 --name minio \
  -v ~/volumes/minio/data:/data \
  -v ~/volumes/minio/config:/root/.minio \
  minio/minio server /datainio/minio server /data

```




#### 参考资料
+ [《minio文档》](https://docs.min.io/)
+ [《基于 MinIO 对象存储框架的短视频点播平台设计》](https://www.163.com/dy/article/GCFV3N2H0511FQO9.html)