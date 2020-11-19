## Go-Zero 快速入门

#### 视频和 PPT

- [《go-zero ppt》](https://docs.google.com/presentation/d/1eRAO2pAjHztrQTpK_2A1CixelFI9Oupht0TkxLTKFG8/edit#slide=id.g8bf5e94d15_1_20)

- [B 站视频介绍 go-zero](https://www.bilibili.com/video/BV1rD4y127PD/)

#### 文档

https://www.yuque.com/tal-tech/go-zero/rm435c

#### 快速使用

```shell
# 创建和启动服务
goctl api new greet
cd greet
go run greet.go -f etc/greet-api.yaml

# 访问接口
curl -i http://localhost:8888/greet/from/you
```

##### 目录结构

```shell
.
├── etc
│   └── greet-api.yaml          # 配置文件
├── greet.api                   # api描述定义问题，如何定义查看文档https://github.com/tal-tech/zero-doc/blob/main/doc/goctl.md
├── greet.go                    # main函数
└── internal
    ├── config
    │   └── config.go           # 配置定义
    ├── handler
    │   ├── greethandler.go     # 路由实现，具体的greet操作由logic包实现
    │   └── routes.go           # 路由定义
    ├── logic
    │   └── greetlogic.go       # 具体的greet内容实现
    ├── svc
    │   └── servicecontext.go   # 定义ServiceContext
    └── types
        └── types.go            # 定义请求体和返回结构
```

#### 源码分析

https://gitee.com/frank2020/go-zero
