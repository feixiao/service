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

### 客户端
#### Golang




#### 参考资料
+ [<<Openfire 的安装和配置>>](https://www.cnblogs.com/hoojo/archive/2012/05/17/2506769.html)
+ [<<OpenFire深入浅出>>](https://pan.baidu.com/s/1hqDxAtQ?errno=0&errmsg=Auth%20Login%20Sucess&&bduss=&ssnerror=0&traceid=)
