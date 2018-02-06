# **go-micro** && micro 

### 快速入门

+ [where-do-i-start](https://micro.mu/docs/faq.html#where-do-i-start)

+ [How do I use Micro?](https://micro.mu/docs/faq.html#how-do-i-use-micro)

+ [Guide](https://micro.mu/docs/install-guide.html)

+ [Micro vs Go-Kit](https://micro.mu/docs/faq.html#micro-vs-go-kit)

  Use go-kit where you want complete control. Use go-micro where you want an opinionated framework.

+ [examples](https://github.com/feixiao/examples)

### Go Micro 

#### 服务发现(Service Discovery)

##### Consul

+ Consul是GoMicro默认使用的服务发现系统，可以使用插件的方式进行修改。可以在  [micro/go-plugins](https://github.com/micro/go-plugins) 查找etcd, kubernetes, zookeeper等系统的插件实现方式。
+ [Consul安装](https://www.consul.io/intro/getting-started/install.html)

##### Multicast DNS

+ [mDNS原理的简单理解](http://www.binkery.com/archives/318.html)

+ [Multicast_DNS](https://en.wikipedia.org/wiki/Multicast_DNS)

+ 使用方式

  ```shell
  go run main.go --registry=mdns
  ```

#### 编写服务(Writing a service)

+ 程序例子在 [examples/service](https://github.com/feixiao/examples/tree/master/service).

#### 编写函数(Writing a Function)

Go Micro 包括支持了函数编程模型。

+ 程序例子在 [examples/function](https://github.com/feixiao/examples/tree/master/function).

#### 插件(Plugins)

默认情况下面go-micro在核心部分对每个接口只提供了很少的一些实现，但是他们完全可以以插件的形式实现。

[可用的插件]( https://github.com/micro/go-plugins)

+ [插件的实现和使用](https://github.com/micro/go-plugins)

#### 包装器(Wrappers)

Go Micro以Wrappers实现中间件的概念。

+ 程序例子在 [service/wrapper](https://github.com/feixiao/examples/tree/master/service/wrapper)

  ​

### Micro 

Micro是微服务开发工具集合，目的是简化分布式系统的开发工作。包括如下五个组成部分：

+ [**Go Micro**](https://micro.mu/docs/go-micro.html) - 开发包：插件式的RPC开发框架。

+ [**API**](https://micro.mu/docs/api.html) -  API网关：单一的Http协议入口，动态的将HTTP请求转换为RPC请求(HTTP请求)。

  ```shell
  micro api --handler=rpc	# 将HTTP请求转换为RPC请求
  micro api # 默认的方式The default handler uses endpoint metadata from the registry to determine service routes
  micro api --handler=proxy  # 使用github.com/micro/go-web，HTTP代理实现
  ```


  + [api example](https://github.com/feixiao/examples/tree/master/api)

+ [**Web**](https://micro.mu/docs/web.html) - 提供了dashboard 查询服务的Web界面。

+ [**CLI**](https://micro.mu/docs/cli.html) - 跟你的微服务交互的命令行接口。

+ [**Bot**](https://micro.mu/docs/bot.html) - The **micro bot** is a bot that sits inside your microservices environment which you can interact with via Slack, HipChat, XMPP, etc. It mimics the functions of the CLI via messaging.

- [**Sidecar**](https://micro.mu/docs/sidecar.html) - go-micro代理。基于HTTP的，go-micro全特性实现。

  ```
  micro sidecar					# rpc 
  micro sidecar --handler=proxy	# http
  ```

- [**New**](https://micro.mu/docs/new.html) - The **micro new** command is a quick way to generate boilerplate templates for micro services.

- [**Run**](https://micro.mu/docs/run.html)  - The **micro run** command manages the lifecycle of a microservice. 

### 参考资料

+ [Micro - a microservices toolkit](https://micro.mu/blog/2016/03/20/micro.html)
+ [Micro Docs](https://micro.mu/docs/index.html)
+ [Micro 架构与设计](https://www.cnblogs.com/s0-0s/p/6874800.html)

