package znet

import (
	"ZinxDemo/ziface"
	"fmt"
	"net"
)

type Connection struct {
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
	// 当前连接的ID 也可以称作为SessionID ID全局唯一
	ConnID uint32
	// 当前连接的关闭状态
	isClosed bool

	// 该连接的处理方法api
	HandleAPI ziface.HandFunc

	// 告知该连接已经退出/停止的channel
	ExitBuffChan chan bool
}

//创建连接的方法
func NewConnection(conn *net.TCPConn, connID uint32, callBackApi ziface.HandFunc) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		HandleAPI:    callBackApi,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}

/* 处理conn读数据的Goroutine */
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running ")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		// 读取我们最大的数据到buf中
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err : ", err)
			c.ExitBuffChan <- true
			continue
		}

		// 调用当前连接业务（这里执行的是当前conn的绑定的handle方法）
		if err := c.HandleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("connID ", c.ConnID, " handle is error ")
			c.ExitBuffChan <- true
			return
		}
	}
}

//启动连接，让当前连接开始工作
func (c *Connection) Start() {
	// 开启处理该连接读取到客户端数据之后的请求业务；
	go c.StartReader()

	for {
		select {
		case <-c.ExitBuffChan:
			// 得到退出消息，不再阻塞
			return
		}
	}
}

//停止连接，结束当前连接状态M
func (c *Connection) Stop() {
	// 1. 如果当前连接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// TODO Connection Stop() 如果用户注册了该链接的关闭回调业务，name在此刻应该显示调用

	// 关闭socket链接
	c.Conn.Close()

	// 通知从缓冲队列读数据的业务，该链接已关闭
	c.ExitBuffChan <- true

	// 关闭该链接全部管道
	close(c.ExitBuffChan)
}

//从当前连接获取原始的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取当前连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
