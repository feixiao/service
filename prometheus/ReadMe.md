# Prometheus

### Prometheus介绍

+ [《基于Prometheus的分布式在线服务监控实践》](https://zhuanlan.zhihu.com/p/24811652?refer=cloudsre)

### 客户机安装exporter

+ 我这边选择[Node/system metrics exporter](https://github.com/prometheus/node_exporter) (**official**),其他工具的选择参考[《Exporters and integrations》](https://prometheus.io/docs/instrumenting/exporters/)

+ node_exporter

  ```
  tar -zxvf node_exporter-0.15.2.linux-amd64.tar.gz
  ./node_exporter
  ```

### Prometheus安装和运行

+  获取安装包 [二进制程序](https://prometheus.io/download/)

+  解压安装

   ```shell
    mv prometheus-2.0.0.linux-amd64.tar.gz  /opt/
    tar -zxvf prometheus-2.0.0.linux-amd64.tar.gz
   ```

+  修改配置文件

   ```shell
    #  添加监控目标
   - job_name: 'node'
      static_configs:
        - targets: ['localhost:9100']  # 注意端口
   ```
+ 启动

  ```
    ./prometheus --config.file=prometheus.yml  
    ./prometheus --config.file=prometheus.yml   --web.listen-address=:8091  # 修改端口
  ```


+ 其他安装方式 [参考资料](https://prometheus.io/docs/prometheus/latest/installation/)
+ 浏览地址： http://172.17.228.13:9090/


### Grafana部署

+ 安装grafana

  ```
  # Ubuntu
  wget https://s3-us-west-2.amazonaws.com/grafana-releases/release/grafana_4.6.3_amd64.deb
  sudo apt-get install -y adduser libfontconfig
  sudo dpkg -i grafana_4.6.3_amd64.deb
  
  # CentOS
  wget https://dl.grafana.com/oss/release/grafana-6.2.5-1.x86_64.rpm
  sudo yum localinstall grafana-6.2.5-1.x86_64.rpm
  ```


+ 启动服务

  ```
  service grafana-server start
  
  systemctl daemon-reload
  systemctl start grafana-server
  systemctl status grafana-server
 
  systemctl enable grafana-server.service
  ```


+ 登入

  http://172.17.228.13:3000  用户名/密码 admin/admin

+ 数据关联

  通过添加数据库的方式进行。
  
+ dashborad选择

  https://grafana.com/dashboards/1860　　（选择13版本）


### 客户机安装exporter

+ 我这边选择[Node/system metrics exporter](https://github.com/prometheus/node_exporter) (**official**),其他工具的选择参考[《Exporters and integrations》](https://prometheus.io/docs/instrumenting/exporters/)

+ node_exporter

  ```
  tar -zxvf node_exporter-0.15.2.linux-amd64.tar.gz
  ./node_exporter
  ```

### 参考资料

+ [《CentOS 7.2下安装Prometheus和Grafana监控MySQL服务器性能》](http://www.linuxidc.com/Linux/2017-06/144972.htm)
+ [《prometheus + grafana安装部署（centos6.8）》](http://www.cnblogs.com/shhnwangjian/p/6878199.html)
+ [《Prometheus 实战》](https://songjiayang.gitbooks.io/prometheus/content/introduction/what.html)
+ [Grafana+Prometheus打造全方位立体监控系统](https://www.cnblogs.com/smallSevens/p/7805842.html)
+ [《通过consul实现prometheus动态服务发现以及告警策略注册》](https://blog.csdn.net/weixin_38645718/article/details/84773305)

