## Traefik快速入门
### 部署基础环境
#### 启动reverse-proxy
```shell
docker-compose up -d reverse-proxy
```

#### 启动whoami
```shell
docker-compose up -d whoami
```

#### 测试
```shell
curl -H Host:whoami.docker.localhost http://127.0.0.1
```

#### 负载均衡测试
```shell
# 启动两个
docker-compose up -d --scale whoami=2

# 测试
curl -H Host:whoami.docker.localhost http://127.0.0.1
```

#### Traefik后台
```shell
http://172.20.99.13:8080/dashboard/#/
```