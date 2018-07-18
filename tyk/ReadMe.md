## tyk快速入门

### 编译

+ 下载tyk源代码 **tyk-2.5.1.zip**

+ 解压到$GOPATH/src/github.com/TykTechnologies/tyk

+ 编译

  ```
  govendor sync
  go build
  ```

### 安装

```
sudo apt-get install redis-server
./tyk # 启动程序
```




### 参考资料

+ [《谈API网关的背景、架构以及落地方案 》](https://mp.weixin.qq.com/s?__biz=MzIwMzg1ODcwMw==&mid=2247486461&amp;idx=1&amp;sn=4f238a1f8046d652f0a05db2bdf55cd7&source=41#wechat_redirect)
+ [《深入浅出聊聊企业级API网关》](https://mp.weixin.qq.com/s?__biz=MzIwMzg1ODcwMw==&mid=2247486230&amp;idx=1&amp;sn=f8d9bba498cb0cebe98e61eb2b7678ba&source=41#wechat_redirect)
+ [《API网关规划与实施细则》](https://www.jianshu.com/p/34cfe5d12793)

