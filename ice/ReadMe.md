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

#### 包含的工具

```shell
turnadmin -> turnserver # coturn的管理工具，主要是生成用户帐户的键值，列出数据库中的可用用户，在数据库中添加和删除用户。
turnserver				# 我们需要的服务器. 
turnutils_natdiscovery
turnutils_oauth
turnutils_peer
turnutils_rfc5769check
turnutils_stunclient		# 用于测试stun服务
turnutils_uclient    		# 测试其TURN服务是否正常。
```



#### 参考资料

+ [stun服务器搭建(coTurn)](https://www.cnblogs.com/wunaozai/p/9071097.html)
+ [webrtc学习: 部署stun和turn服务器](https://www.cnblogs.com/lidabo/p/6971103.html)

