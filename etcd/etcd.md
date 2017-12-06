## Etcd学习整理

Etcd是用于共享配置和服务发现的分布式，一致性的KV存储系统。etcd作为一个受到ZooKeeper与doozer启发而催生的项目，除了拥有与之类似的功能外，更专注于以下四点。

- 简单：基于HTTP+JSON的API让你用curl就可以轻松使用。
- 安全：可选SSL客户认证机制。
- 快速：每个实例每秒支持一千次写操作。
- 可信：使用Raft算法充分实现了分布式。

### Raft实现

+ [etcd/raft](https://github.com/coreos/etcd/tree/master/raft)

### 参考资料

+ [《ETCD系列一：简介》](https://yq.aliyun.com/articles/11035?spm=5176.8091938.0.0.v9RtUj)
+ [《etcd：从应用场景到实现原理的全方位解读》](http://www.infoq.com/cn/articles/etcd-interpretation-application-scenario-implement-principle)
+ [《CoreOS 实战：剖析 etcd》](http://www.infoq.com/cn/articles/coreos-analyse-etcd/)