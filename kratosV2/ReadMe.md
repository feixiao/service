## Kratos V2熟悉

#### 安装
```shell
go env -w GO111MODULE=on
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

# 验证版本
➜  ~ kratos -v
kratos version v2.0.0-rc1
```


#### 使用
```shell
# 创建项目模板
kratos new helloworld

cd helloworld
# 拉取项目依赖
go mod download

# 生成proto模板
kratos proto add api/helloworld/helloworld.proto

# 生成proto源码
kratos proto client api/helloworld/helloworld.proto

# 生成server模板
kratos proto server api/helloworld/helloworld.proto -t internal/service
```