## 本地搭建 Istio 环境

### Minikube

#### 安装

```shell
curl -Lo minikube http://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/releases/v1.16.0/minikube-linux-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/
```

#### 启动

```shell
#   --insecure-registry="172.18.0.57:5000"
minikube start --vm-driver=docker --registry-mirror=https://registry.docker-cn.com --image-mirror-country='cn' --image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers' --kubernetes-version v1.18.0 --cpus=4 --memory=8192
```

#### kubectl

```shell
curl -Lo kubectl http://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/linux/amd64/kubectl && chmod +x ./kubectl && sudo mv ./kubectl /usr/local/bin/kubectl
```
