## BookInfo部署

#### 创建新的namespace
```shell
kubectl create ns bookinfo

kubectl label namespace bookinfo istio-injection=enabled
```

#### 部署服务
```shell
# 部署服务
kubectl apply -f ./bookinfo/platform/kube/bookinfo.yaml -n bookinfo

# 查看服务
[frank@node samples]$ kubectl get service -n bookinfo
NAME          TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
details       ClusterIP   10.105.147.235   <none>        9080/TCP   72s
productpage   ClusterIP   10.107.73.208    <none>        9080/TCP   72s
ratings       ClusterIP   10.106.13.108    <none>        9080/TCP   72s
reviews       ClusterIP   10.101.90.80     <none>        9080/TCP   72s

# 查看pod
[frank@node samples]$ kubectl get pods -n bookinfo
details-v1-558b8b4b76-l56z9       2/2     Running   0          66m
productpage-v1-6987489c74-9lxdh   2/2     Running   0          66m
ratings-v1-7dc98c7588-tq662       2/2     Running   0          66m
reviews-v1-7f99cc4496-6klkl       2/2     Running   0          66m
reviews-v2-7d79d5bd5d-9svct       2/2     Running   0          66m
reviews-v3-7dbcdcbc56-65qg2       2/2     Running   0          66m
```

#### 部署Gateway
```shell
kubectl apply -f ./bookinfo/networking/bookinfo-gateway.yaml  -n bookinfo
kubectl get gateway -n bookinfo
```

#### 设置
```shell
kubectl get svc istio-ingressgateway -n istio-system

# for minikube
export INGRESS_HOST=$(minikube ip)
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT
```


### 参考资料
+ [《examples/bookinfo》](https://istio.io/latest/docs/examples/bookinfo/)