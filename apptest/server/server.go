package main

import (
	"ZinxDemo/ziface"
	"ZinxDemo/znet"
	"fmt"
)

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter // 一定要先基础BaseRouter
}

//Test PreHandle
//func (pr *PingRouter) PreHandle(request ziface.IRequest) {
//	fmt.Println("Call Router PreHandle")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ...\n"))
//	if err != nil {
//		fmt.Println("call back ping ping ping error")
//	}
//}

//Test Handle
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	// 先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	// 会写数据
	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

//Test PostHandle
//func (pr *PingRouter) PostHandle(request ziface.IRequest) {
//	fmt.Println("Call Router PostHandle")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
//	if err != nil {
//		fmt.Println("call back ping ping ping error")
//	}
//}

// Server 模块的测试函数
func main() {
	// 1 创建一个 server 句柄 s
	s := znet.NewServer("[ zinx v0.3]")

	s.AddRouter(&PingRouter{})

	// 2 开启服务
	s.Server()
}
