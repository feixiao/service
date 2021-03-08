## 本地搭建 Istio 环境

### Minikube

#### 安装

```shell
curl -Lo minikube http://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/releases/v1.16.0/minikube-linux-amd64 &&  chmod +x minikube && sudo mv minikube /usr/local/bin/
```

#### 启动

```shell
#   --insecure-registry="172.18.0.57:5000"
minikube start --vm-driver=docker --registry-mirror=https://registry.docker-cn.com --image-mirror-country='cn' --image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers' --kubernetes-version v1.18.0 --cpus=4 --memory=8192
```

#### kubectl

```shell
curl -Lo kubectl http://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/linux/amd64/kubectl && chmod +x ./kubectl && sudo mv ./kubectl /usr/local/bin/kubectl

# 验证
kubectl cluster-info
```

### Helm

```shell
wget https://get.helm.sh/helm-v3.5.0-linux-amd64.tar.gz
tar -zxvf helm-v3.5.0-linux-amd64.tar.gz
sudo mv linux-amd64/helm /usr/local/bin/helm
```

### Istio

```shell
# 下载页面
https://github.com/istio/istio/releases/tag/1.8.2
tar -zxvf istio-1.8.2-linux-amd64.tar.gz
sudo mv istio-1.8.2 /opt/

# 配置PATH环境变量

# 然后验证
istioctl install --set profile=demo -y
```

#### 案例

- [《Istio 服务网格进阶实战》](https://www.servicemesher.com/istio-handbook/)
- [《cloud-native-istio》](https://github.com/feixiao/cloud-native-istio)
