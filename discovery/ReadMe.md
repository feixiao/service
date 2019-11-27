## discovery快速入门
### 简介
discovery是B站参考Eureka自研的服务注册发现系统。其前期使用Zookeeper为主，但是由于遇到如下一系列问题，从而自研discovery。

+ 用好难，维护难，二次开发难
+ 服务规模太大(3W+节点)，一套zk集群很难满足大量的连接和写请求，但如果拆分服务类型维护多套集群又破坏了连通性
+ CP系统，对于微服务的服务注册发现，其实不如一套AP系统更可用

### 安装部署
+ WSL 不能编译运行（无法采集系统基础信息）
+ Windows上使用Docker
    ```
    docker pull golang:1.12.7
    docker run -it --name gobuild -v D:\volumes\golang:/go golang

    # 配置golang变量
    export GOPROXY=https://goproxy.io
    export GO111MODULE=on
    ```
+ 编译运行
    ```
    cd discovery/cmd/discovery
    go build
    ./discovery -conf discovery-example.toml
    ```

### 客户端接入(参考demo)

### 参考资料
+ [discovery介绍](https://github.com/bilibili/discovery/blob/master/doc/intro.md)
+ [设计概述](https://github.com/bilibili/discovery/blob/master/doc/arch.md)
+ [最佳事件](https://github.com/bilibili/discovery/blob/master/doc/practice.md)