## gRPC快速入门

#### 预备知识

+ gRPC需要Go 1.6版本以上

  ```shell
  $ go version
  go version go1.12 linux/amd64   # 使用的ubuntu上面的go1.8版本
  ```

#### 安装gRPC

  + 从github获取最新的开发版本

    ```shell
    go get -u google.golang.org/grpc  # 需要翻墙(go mod以后配置GOPROXY以后就不需要了)
    ```

+ 安装protobuf

  ```shell
  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

  # 为了安装命令行工具protoc
  https://github.com/google/protobuf/releases/tag/v3.4.0
  ```


#### 例子分析
+ 使用流程
  + 在一个 .proto 文件内定义服务。
  + 用 protocol buffer 编译器生成服务器和客户端代码。
  + 使用 gRPC 的 Go API 为你的服务实现一个简单的客户端和服务器。

+ 项目结构

  ```shell
    .
    ├── client
    │   └── client.go
    ├── mock_routeguide
    │   ├── rg_mock.go
    │   └── rg_mock_test.go
    ├── routeguide
    │   └── route_guide.proto　　＃　定义服务
    ├── server
    │   └── server.go
    └── testdata
      └── route_guide_db.json
  ```

#### 参考资料
+ [《gRPC 官方文档中文版》](http://doc.oschina.net/grpc?t=56831)
+ [Go Quick Start](https://grpc.io/docs/quickstart/go.html)
+ [protocol-buffers](https://developers.google.com/protocol-buffers/)
+ [gRPC初接触](https://samael65535.github.io/2017-05-18/grpc_newb/)
+ [Go gRPC 调试工具 ](http://securedsearch.lavasoft.com/)
+ [Grpc兼容HTTP协议文档自动生成网关](https://my.oschina.net/wenzhenxi/blog/3023874)