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

- [great/ReadMe.md](./greet/ReadMe.md)

#### 官方完整的 Demo

- [bookstore](https://github.com/tal-tech/go-zero/tree/master/example/bookstore)
- [文档](https://www.yuque.com/tal-tech/go-zero/rm435c)

#### 源码分析

https://gitee.com/frank2020/go-zero
