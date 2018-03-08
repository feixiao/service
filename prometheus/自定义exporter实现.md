## 自定义exporter实现

### [Concepts](https://prometheus.io/docs/concepts/data_model/#)

#### [Data model](https://prometheus.io/docs/concepts/data_model/)

**prometheus** 原理上以[Time_series](https://en.wikipedia.org/wiki/Time_series) 的方式存储数据。流的时间戳值属于同样的度量和同样的标记维度。

+ Metric names and labels                                                                                                                                                                   

  + **metric name** 指定了系统的一般特性,比如http_request_total 表示全部接收到的HTTP请求。
  + Label启用了Prometheus的多维数据模型，对于相同的**metric name**任何给定的标签组合指示这个度量的一个特定维度的实例化。

+ Notation

  + 给定一个**metric name**和一组**Label**，时间序列经常使用这个符号来标识

    ```
    <metric name>{<label name>=<label value>, ...}
    ```

  +  举个例子，比如一个时间序列,**metric name**为api_http_requests_total, label名字为method=“POST”,同时handler="/messages",可以写成下面的格式。

    ```
    api_http_requests_total{method="POST", handler="/messages"}
    ```


#### [Metric types](https://prometheus.io/docs/concepts/metric_types/)

+ **Counter**

  counter表示一个累计的度量，该值只会增加。


+ **Gauge**

  gauge表示单个数值的度量上可以任意增减。

+ **Histogram**

  Histogram表示样例观察，并将它们放在可配置的桶（configurable buckets）中。它还提供了所有观测值的总和。

+ **Summary**

  Summary跟Histogram类似。它还提供了一个观测的总数量和所有观测值的总和,它计算可配置的分位数滑动时间窗口。

### Sample

[examples/random](https://github.com/prometheus/client_golang/tree/master/examples/random) 观察数据采集的效果。

### [Writing exporters](https://prometheus.io/docs/instrumenting/writing_exporters/)

#### 官方建议:

+ 如果系统的度量指标很少变化，可以参考[haproxy_exporter](https://github.com/prometheus/haproxy_exporter) 。
+ 如果系统的度量指标在新的版本中经常变化，可以参考[mysqld_exporter](https://github.com/prometheus/mysqld_exporter)。
+ [node_exporter](https://github.com/prometheus/node_exporter)是这些组件的混合，其复杂性随模块的不同而变化。

### 参考资料

+ [《Prometheus监控系统小结》](https://weibo.com/p/23041882419ac10102x1ah?sudaref=www.baidu.com&display=0&retcode=6102&sudaref=passport.weibo.com)
+ [官方提供的exporters ](https://prometheus.io/docs/instrumenting/exporters/)
+ [writing_exporters](https://prometheus.io/docs/instrumenting/writing_exporters/)

