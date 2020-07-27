## Traefik快速入门

### 使用docker部署
```shell
# 准备配置文件
cd /home/frank/volumes/traefik

wget https://raw.githubusercontent.com/containous/traefik/v2.2/traefik.sample.toml -o traefik.toml

docker run -d -p 8080:8080 -p 80:80 --name="traefik" \
    -v /home/frank/volumes/traefik/traefik.toml:/etc/traefik/traefik.toml traefik:v2.2
```



### 参考资料
+ [《Traefik简易入门》](https://shanyue.tech/op/traefik.html#%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B)