## Grafana使用记录


#### Auth Proxy
+ 运行
```shell
# 准备配置grafana.ini，运行grafana
docker run --rm -i -p 3000:3000 -v $(pwd)/auth/grafana.ini:/etc/grafana/grafana.ini --name grafana grafana/grafana:7.2.0-ubuntu
```
+ 获取用户信息
```shell
curl -H "X-WEBAUTH-USER: admin"  http://localhost:3000/api/users
[{
	"id": 1,
	"name": "",
	"login": "admin",
	"email": "admin@localhost",
	"avatarUrl": "/avatar/46d229b033af06a191ff2267bca9ae56",
	"isAdmin": true,
	"isDisabled": false,
	"lastSeenAt": "2020-10-14T02:07:04Z",
	"lastSeenAtAge": "\u003c 1m",
	"authLabels": ["OAuth"]
}]
```
+ 运行httpd
```shell
# 准备配置httpd.conf
# 添加用户名 frank 密码 1
sudo apt-get install apache2-utils
htpasswd -bc htpasswd frank 1
# link到grafana,所以配置文件里面的hostname grafana可以被解析
docker run --rm -i -p 80:80 --link grafana:grafana -v $(pwd)/auth/httpd.conf:/usr/local/apache2/conf/httpd.conf -v $(pwd)/auth/htpasswd:/tmp/htpasswd httpd:2.4

# 浏览器访问
http://{ip}

# 输入密码， 我们既可以访问grafana
```