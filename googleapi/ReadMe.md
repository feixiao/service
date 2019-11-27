## Google API快速入门

### 简介
使用Protocol Buffers定义RPC API, 这些API以gRPC的方式发布，同时可以通过gRPC-HTTP转换可以作为REST API使用。

### API设计规范
+ [《API-design-guide中文版》](https://github.com/feixiao/gitbook/tree/master/API-design-guide)


### 如何使用
+ [谷歌api的公共接口定义](https://github.com/googleapis/googleapis)
##### 运行例子
+ 获取代码
    ```shell
    git clone https://github.com/googleapis/googleapis.git    
    ```
+ 安装Bazel
    ```
    # ubuntu 18.04
    sudo apt-get install openjdk-11-jdk
    sudo apt-get install curl
    curl https://bazel.build/bazel-release.pub.gpg | sudo apt-key add -
    echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" \
    | sudo tee /etc/apt/sources.list.d/bazel.list

    sudo apt-get update && sudo apt-get install bazel 
    ```
+ 安装库()
    ```
    bazel build //google/example/library/v1/...

    # java 
    bazel build //google/example/library/v1:google-cloud-library-v1-java
    ```
+ Go语言
    ```
    https://github.com/googleapis/go-genproto

    go get -u github.com/golang/protobuf/protoc-gen-go 
    
    ```
### 参考资料
+ [官网](https://googleapis.github.io/)
