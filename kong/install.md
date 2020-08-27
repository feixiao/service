## Kong安装

#### [Docker Installation](https://docs.konghq.com/install/docker/)

```shell
# 获取镜像
docker pull kong:2.1.3-centos
docker pull postgres:9.6

# 创建docker网络，用于容器之间的通信
docker network create kong-net

#创建kong使用的数据库
docker run -d --name kong-database \
              --network=kong-net \
              -p 5432:5432 \
              -e "POSTGRES_USER=kong" \
              -e "POSTGRES_DB=kong" \
              -e "POSTGRES_PASSWORD=kong" \
              postgres:9.6

# 配置数据库
docker run --rm \
     --network=kong-net \
     -e "KONG_DATABASE=postgres" \
     -e "KONG_PG_HOST=kong-database" \
     -e "KONG_PG_USER=kong" \
     -e "KONG_PG_PASSWORD=kong" \
     -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
     kong:2.1.3-centos kong migrations bootstrap
  
    
#启动Kong           
docker run -d --name kong \
     --network=kong-net \
     -e "KONG_DATABASE=postgres" \
     -e "KONG_PG_HOST=kong-database" \
     -e "KONG_PG_USER=kong" \
     -e "KONG_PG_PASSWORD=kong" \
     -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
     -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
     -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
     -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
     -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
     -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" \
     -p 8100:8000 \
     -p 8443:8443 \
     -p 8001:8001 \
     -p 8444:8444 \
     kong:2.1.3-centos
    
    
# 浏览器确认
http://172.20.99.13:8001/
```

#### 参考资料
+ [《Docker Installation》](https://docs.konghq.com/install/docker/)
