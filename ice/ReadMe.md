## coturn快速入门

#### 编译安装

```shell
wget https://github.com/downloads/libevent/libevent/libevent-2.0.21-stable.tar.gz
tar xvfz libevent-2.0.21-stable.tar.gz
cd libevent-2.0.21-stable
./configure
make
make install

tar -zxvf coturn-4.5.0.8.tar.gz 
cd coturn-4.5.0.8
./configure 
make
make install

# 配置文件在/usr/local/etc
```



#### 参考资料

+ [stun服务器搭建(coTurn)](https://www.cnblogs.com/wunaozai/p/9071097.html)