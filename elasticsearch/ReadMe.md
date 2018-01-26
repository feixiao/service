## elasticsearch

### 安装和运行

```shell  
tar -zxvf elasticsearch-6.1.1.tar.gz 
cd elasticsearch-6.1.1/bin
./elasticsearch -d  # 不要使用root用户启动

# 查看服务是否启动成功
netstat -apn | grep 9200

```

+ [ES启用外部IP访问的启动异常](http://kael-aiur.com/elk/ES%E9%85%8D%E7%BD%AE%E7%BB%99%E5%A4%96%E9%83%A8%E6%9C%BA%E5%99%A8%E9%80%9A%E8%BF%87ip%E8%AE%BF%E9%97%AE.html)

### 加载测试数据

+ [Loading Sample Data](https://www.elastic.co/guide/en/kibana/current/tutorial-load-dataset.html)

### 安装**elasticsearch-sql**

```sh
./bin/elasticsearch-plugin install https://github.com/NLPchina/elasticsearch-sql/releases/download/6.1.1.0/elasticsearch-sql-6.1.1.0.zip
```

访问

```
http://172.17.228.97:9200/_plugin/sql/
```

### elasticsearch-head

```
sudo apt-get install nodejs-legacy
npm install
git clone git://github.com/mobz/elasticsearch-head.git
cd elasticsearch-head
npm install
npm run start
```

访问

```
http://172.17.228.97:9100
```

