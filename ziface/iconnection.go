package ziface

import "net"

type IConnection interface {
	// 启动连接，让当前连接开始工作
	Start()
	// 停止连接，结束当前连接状态
	Stop()
	// 从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	// 获取当前连接ID
	GetConnID() uint32
	// 获取远程客户端地址信息
	RemoteAddr() net.Addr

	// 直接将 Message 数据发送数据给远程的 TCP 客户端
	SendMsg(msgId uint32, data []byte) error
}

// 定义一个统一处理链接业务的接口
// 第一参数是socket原生链接
// 第二个参数是客户端请求的数据
// 第三个参数是客户端请求的数据长度
// 如果我们想要指定一个conn的处理业务，只要定义一个HandFunc类型的函数，然后和该链接绑定就可以了
//type HandFunc func(*net.TCPConn, []byte, int) error
