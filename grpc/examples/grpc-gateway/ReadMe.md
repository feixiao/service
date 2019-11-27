## grpc-gateway入门

#### 生成文件          
```shell
// 生成grpc结构文件
protoc --proto_path=./ -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. gateway.proto

// 生成gateway文件
protoc --proto_path=../ -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. gateway.proto
```


#### 参考资料
+ [《Grpc-Gateway - Grpc兼容HTTP协议文档自动生成网关》](https://my.oschina.net/wenzhenxi/blog/3023874)