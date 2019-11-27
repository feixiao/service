package main

import (
	"fmt"
	"time"
	"github.com/bilibili/discovery/naming"
	"os"
	"os/signal"
	"syscall"
	"log"
)


func main() {
	conf := &naming.Config{
		Nodes: []string{"127.0.0.1:10001"}, // NOTE: 配置种子节点(1个或多个)，client内部可根据/discovery/nodes节点获取全部node(方便后面增减节点)
		Zone:  "sh001",  // 字段需要跟服务器匹配
		Env:   "dev",
	}
	dis := naming.New(conf)
	ins := &naming.Instance{
		Zone:  "sh001",	
		Env:   "dev",
		AppID: "provider",
		// Hostname:"", // NOTE: hostname 不需要，会优先使用discovery new时Config配置的值，如没有则从os.Hostname方法获取！！！
		Addrs:    []string{"http://172.0.0.1:8888", "grpc://172.0.0.1:9999"},
		LastTs:   time.Now().Unix(),
		Metadata: map[string]string{"weight": "10"},
	}
	cancel, _ := dis.Register(ins)
	defer cancel() // NOTE: 注意一般在进程退出的时候执行，会调用discovery的cancel接口，使实例从discovery移除
	fmt.Println("register")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Println("discovery get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			cancel()
			time.Sleep(time.Second)
			log.Println("discovery exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}