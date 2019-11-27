# HTTP 隧道代理原理和实现

### 服务端
+ [《ngx_http_proxy_connect_module》](https://github.com/chobits/ngx_http_proxy_connect_module)
+ [charlesproxy](https://www.charlesproxy.com/)
+ [mitmproxy](https://mitmproxy.org/)
+ [HTTPTunnel](https://github.com/larsbrinkhoff/httptunnel)
    + [HTTPTunnel Package Description](https://tools.kali.org/maintaining-access/httptunnel)
    ```
    # 172.17.229.3
    nc -l 10000
    # 转发数据(20000端口到数据转发到172.17.229.3:10000)
    hts  -w -F localhost:10000 20000

    # 172.17.229.2
    # 本地端口 8090 转发到 172.17.229.3:10000
    htc -T 30000 -k 20000 -w -F 8090 172.17.229.3:20000
  
    telnet localhost 8090  // 正常情况应该转发到172.17.229.3:10000
    ```
### 参考资料
+ [《HTTP 隧道代理原理和实现》](https://cizixs.com/2017/03/22/http-tunnel-proxy-and-golang-implementation/)
+ [《什么是HTTP隧道，怎么理解HTTP隧道呢？》](https://www.zhihu.com/question/21955083)
