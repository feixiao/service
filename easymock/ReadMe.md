## EasyMock快速入门

#### 服务部署
```shell
git clone https://github.com/easy-mock/easy-mock-docker.git
cd easy-mock-docker

# 修改日志路径权限
chmod 777 /home/*/logs

# 根据说明配置好配置文件
docker-compose up -d
```

#### 登入服务
```
# 浏览器浏览 http://localhost:7300/
默认用户名为root 密码为123456
```

#### 参考资料
+ [《使用 docker 运行 easy-mock》](https://zhuanlan.zhihu.com/p/91646508)