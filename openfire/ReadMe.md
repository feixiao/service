## OpenFire快速入门                                                                                                                               
#### XMPP介绍                                                                                                                            

### 服务器端
+ 获取镜像
    ```
    docker pull gizmotronic/openfire:4.3.2
    ```

+ Ubuntu
    ```
    docker run --name openfire -d --restart=always \
        --publish 9090:9090 --publish 5222:5222 --publish 7777:7777 \
        --volume /home/frank/docker/openfire:/var/lib/openfire \
        gizmotronic/openfire:4.3.2
    ```
+ Windows
    ```
    docker run --name openfire -d --restart=always  --publish 9090:9090     --publish 5222:5222 --publish 7777:7777   --volume D:/docker_volume/    openfire:/var/lib/openfire gizmotronic/openfire:4.3.2
    ```



+ 管理页面 http://localhost:9090  注意修改 域


### 客户端
+ [go-xmpp](https://github.com/FluuxIO/go-xmpp) Golang客户端库
+ [gloox](https://camaya.net/gloox/) C++客户端库
+ [poezio](https://github.com/poezio/poezio) Python命令行客户端程序
+ [slixmpp](https://github.com/poezio/slixmpp) Python客户端库
+ [spark](https://github.com/igniterealtime/Spark) 官方客户端程序


#### 参考资料
+ [《XMPP》](https://xmpp.org/)
+ [《Openfire 的安装和配置》](https://www.cnblogs.com/hoojo/archive/2012/05/17/2506769.html)
+ [《OpenFire深入浅出》](https://pan.baidu.com/s/1hqDxAtQ?errno=0&errmsg=Auth%20Login%20Sucess&&bduss=&ssnerror=0&traceid=)
