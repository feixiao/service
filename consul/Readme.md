# consul 入门

### Consul介绍

Consul是一个服务管理软件，支持多数据中心下，分布式高可用的，服务发现和配置共享。

主要特性如下：

+ 服务发现
+ 健康检查
+ 键值对存储
+ 多数据中心

### [Consul vs. Other Software](https://www.consul.io/intro/vs/index.html)

### 快速开始

#### 安装

+ 下载二进制[安装包](https://www.consul.io/downloads.html)

+ 添加到PATH路径

  ```shell
  sudo mv consul /usr/local/bin/
  ```

#### 运行agent

​	Agent可以运行在Server模式或者Client模式。每个数据中心至少有一个Server。其他节点运行Client模式。Client模式是一个非常轻量级的进程，它们注册服务，执行健康检查，将查询转发到服务器。Agent必须运行在集群的全部节点。

+ 以开发模式运行节点(不建议在生产环境使用)

  ```shell
  consul agent -dev
  ```

  上面consul agent启动在控制台输出一些日志，从日志我们可以知道Agent运行在Server模式，同时已经选举自己为集群的“领导”。

+ 查看集群的成员

  ```shell
  ➜  ~ consul members
  Node      Address         Status  Type    Build  Protocol  DC   Segment
  frank-PC  127.0.0.1:8301  alive   server  1.0.1  2         dc1  <all>
  ```

  注：元数据通过添加-detailed参数获取。

  + members命令基于gossip协议，结果最终一致性。如果想获取强一致性的结果，可以通过HTTP API获取。

    ```
    curl localhost:8500/v1/catalog/nodes
    ```

  + 此外，[DNS interface](https://www.consul.io/docs/agent/dns.html)可以被用来查询节点。

+ 停止Agent

​	使用 `Ctrl-C` 友好的退出。

#### 服务

##### 定义服务

注册服务可以通过提供[服务定义](https://www.consul.io/docs/agent/services.html)或者发送[HTTP API](https://www.consul.io/api/index.html)完成。

+ 提供服务定义是比较常见的方法，所以我们通过这种方式进行演示。

  ```shell
  #在当前目录创建consul.d文件夹
  mkdir consul.d
  echo '{"service": {"name": "web", "tags": ["rails"], "port": 80}}' \
  	>     | sudo tee ./consul.d/web.json
  ```


+ 通过上面的配置文件启动agent

  ```shell
  consul agent  -dev -config-dir=./consul.d
  ```

  当log显示agent: Synced service "web"，说明agent成功从配置文件中获取服务定义，然后成功地注册到Service Catalog(下面服务查询中使用的地址);如果想注册多个服务在目录下面配置多个服务定义文件即可。

##### 查询服务

一旦Agent启动，服务器被同步，我们可以通过HTTP API或者DNS查询。

+ DNS API

  服务的默认DNS名字为 **NAME.service.consul**，

  ```
  dig @127.0.0.1 -p 8600 web.service.consul
  ```

+ HTTP API

  ```
  curl http://localhost:8500/v1/catalog/service/web
  ```

##### 更新服务

可以修改文件然后发送SGIHUP信号给Agent修改。或者通过HTTP API进行修改。

#### Consul集群

#### 健康检查

#### 键值数据

#### Web UI

### 参考资料

+ [ 《intro》](https://www.consul.io/intro/index.html)


+ [ 《 consul 入门》](http://blog.csdn.net/viewcode/article/details/45915179)