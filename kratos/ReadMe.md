## Kratos快速入门
### 简介
Kratos是bilibili开源的一套Go微服务框架，包含大量微服务相关框架及工具。

### 安装使用
#### 依赖

```
Go version>=1.12 and GO111MODULE=on
```

#### 安装

```shell
go get -u github.com/bilibili/kratos/tool/kratos
cd $GOPATH/src
kratos new kratos-demo

#　带grpc的项目
kratos new kratos-demo --proto
```

#### 运行

```shell
cd kratos-demo/cmd
go build
./cmd -conf ../configs
```

### 详细介绍

#### http blademaster

+ 内置中间件

  + Recovery

    用于recovery panic。

  + Trace

    用于trace设置，并且实现了net/http/httptrace的接口，能够收集官方库内的调用栈详情。

  + Logger

    用于请求日志记录。

  + CSRF

    用于防跨站请求。

  + CORS

    用于跨域允许请求。

  + 自适应限流

    [自适应限流保护](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/ratelimit.md)

+ 默认的路由

  ```shell
  # 查看默认路由
  curl 'http://127.0.0.1:8000/metadata'
  ```


+ 自动生成代码和文档

  [快速生成bm框架对应的代码](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/blademaster-pb.md)

#### grpc warden

+ GRPC的简单封装，集成trace、log、prom组件　[《quickstart》](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/warden-quickstart.md)

+ 可选择注册发现系统，zookeeper、etcd　[《服务发现》](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/warden-resolver.md)

+ [负载均衡](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/warden-balancer.md)

+ [拦截器](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/warden-mid.md)

  

#### config

+ [本地配置](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/config.md)
+ 配置中心使用携程的apollo

#### ecode

　+ [错误码](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/ecode.md)

#### trace 缺失文档

#### log

+ [日志基础库](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/config.md)
+ log-agent目前　缺失文档

#### database

+ 数据库驱动，进行封装加入了熔断、链路追踪和统计，以及链路超时。
+ [详细使用](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/database.md)

#### cache

+ 统一的cache包，用于进行各类缓存操作。
+ [redis/memcache](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/cache.md)

#### kratos工具

+  [protoc](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/kratos-protoc.md)　　protoc命令的封装。

+  [swagger](kratos-swagger.md)　自动生成swagger文档。

+  [genmc](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/kratos-genmc.md)　缓存代码生成。

+  [genbts](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/kratos-genbts.md)  缓存回源代码生成。

#### 限流bbr

+ 自动嗅探负载和 qps，减少人工配置自动嗅探负载和 qps，减少人工配置
+ 削顶，保证超载时系统不被拖垮，并能以高水位 qps 继续运行
+ [bbr](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/ratelimit.md)

#### 熔断breaker

[熔断器](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/breaker.md)是为了当依赖的服务已经出现故障时，主动阻止对依赖服务的请求。保证自身服务的正常运行不受依赖服务影响，防止雪崩效应。

+ 内置breaker的组件

      RPC client: pkg/net/rpc/warden/client
      Mysql client：pkg/database/sql
      Tidb client：pkg/database/tidb
      Http client：pkg/net/http/blademaster

#### UT单元测试



#### 参考资料

+ [kratos](https://github.com/bilibili/kratos)
+ [kratos使用介绍](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/kratos-tool.md)
+ [文档索引](https://github.com/bilibili/kratos/blob/master/doc/wiki-cn/summary.md)

  