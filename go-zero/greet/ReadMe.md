## go-zero 框架快速入门

#### 目录结构

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
