# Prometheus

### Prometheus介绍

+ [《基于Prometheus的分布式在线服务监控实践》](https://zhuanlan.zhihu.com/p/24811652?refer=cloudsre)

### Prometheus安装和运行

+ 获取安装包 [二进制程序](https://prometheus.io/download/)

+ 解压安装

  ```shell
   mv prometheus-2.0.0.linux-amd64.tar.gz  /opt/
   tar -zxvf prometheus-2.0.0.linux-amd64.tar.gz
  ```

+ 启动

  ```shell
  ./prometheus --config.file=prometheus.yml
  ```

+ 其他安装方式 [参考资料](https://prometheus.io/docs/prometheus/latest/installation/)


### Grafana部署

+ 安装grafana

  ```
  wget https://s3-us-west-2.amazonaws.com/grafana-releases/release/grafana_4.6.3_amd64.deb
  sudo apt-get install -y adduser libfontconfig
  sudo dpkg -i grafana_4.6.3_amd64.deb
  ```


+ 启动服务

  ```
  service grafana-server start
  ```


+ 登入

  http://172.17.228.13:3000  用户名/密码 admin/admin

+ 数据关联

  通过添加数据库的方式进行。

### MySQL的dashboards

+ 下载[grafana-dashboards](https://github.com/percona/grafana-dashboards/releases/tag/v1.5.3)

+ 安装和重启grafana

  ```shell
  cp -rf grafana-dashboards-1.5.3/dashboards /var/lib/grafana/dashboards

  #修改配置文件
  vi /etc/grafana/grafana.ini
  [dashboards.json]
  enabled = false
  path = /var/lib/grafana/dashboards

  #重启
  /etc/init.d/grafana-server restart
  ```

### mysqld_exporter && node_exporter

+ node_exporter

  ```
  tar -zxvf node_exporter-0.15.2.linux-amd64.tar.gz
  ./node_exporter
  ```







### 客户机安装exporter

+ 我这边选择[Node/system metrics exporter](https://github.com/prometheus/node_exporter) (**official**),其他工具的选择参考[《Exporters and integrations》](https://prometheus.io/docs/instrumenting/exporters/)
+ ​

### 参考资料

+ [《CentOS 7.2下安装Prometheus和Grafana监控MySQL服务器性能》](http://www.linuxidc.com/Linux/2017-06/144972.htm)
+ [《prometheus + grafana安装部署（centos6.8）》](http://www.cnblogs.com/shhnwangjian/p/6878199.html)

