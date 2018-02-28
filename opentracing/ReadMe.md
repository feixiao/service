# **OpenTracing**

### OpenTracing

OpenTracing通过提供平台无关、厂商无关的API，使得开发人员能够方便的添加（或更换）追踪系统的实现。  

+ [appdash](https://github.com/sourcegraph/appdash)
+ [zipkin](https://zipkin.io/)
+ [lightstep](https://lightstep.com/)

#### OpenTracing快速入门

+ [《opentracing文档中文版》](https://www.gitbook.com/book/wu-sheng/opentracing-io/details)

### Go工具包

+ [go-kit/kit/tracing ](https://github.com/go-kit/kit/tree/master/tracing)  `package tracing` provides [Dapper](http://research.google.com/pubs/pub36356.html)-style request tracing to services.
+ [《 Tracing HTTP request latency in Go with OpenTracing》](https://medium.com/opentracing/tracing-http-request-latency-in-go-with-opentracing-7cc1282a100a) 客户端使用[**net/http/httptrace**](https://golang.org/pkg/net/http/httptrace/) 集成opentracing。
+ [opentracing-contrib/go-stdlib](https://github.com/opentracing-contrib/go-stdlib) 服务端使用opentracing中间件。

### 参考资料

+ [[Zipkin快速开始](https://segmentfault.com/a/1190000012342007)](https://segmentfault.com/a/1190000012342007)
+ [《opentracing文档中文版》](https://www.gitbook.com/book/wu-sheng/opentracing-io/details)


