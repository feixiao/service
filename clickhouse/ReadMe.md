## ClickHouse 快速入门

#### 安装
```shell
sudo apt-get install apt-transport-https ca-certificates dirmngr
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv E0C56BD4

echo "deb https://repo.clickhouse.com/deb/stable/ main/" | sudo tee \
    /etc/apt/sources.list.d/clickhouse.list
sudo apt-get update

sudo apt-get install -y clickhouse-server clickhouse-client  # 密码设置了123456

# 服务端
sudo service clickhouse-server start

# 客户端
clickhouse-client # or "clickhouse-client --password" if you set up a password.
```


#### 参考资料
+ [https://clickhouse.com/#quick-start](https://clickhouse.com/#quick-start)