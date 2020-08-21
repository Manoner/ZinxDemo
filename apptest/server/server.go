package main

import (
	"ZinxDemo/ziface"
	"ZinxDemo/znet"
	"fmt"
)

type PingRouter struct {
	znet.BaseRouter // 一定要先基础BaseRouter
}

//Test PreHandle
func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ...\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

//Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Server 模块的测试函数
func main() {
	// 1 创建一个 server 句柄 s
	s := znet.NewServer("[ zinx v0.3]")

	s.AddRouter(&PingRouter{})

	// 2 开启服务
	s.Server()
}
