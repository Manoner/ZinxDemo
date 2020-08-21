package znet

import (
	"ZinxDemo/ziface"
	"errors"
	"fmt"
	"net"
	"time"
)

// iServer 接口实现，定义一个Server服务类
type Server struct {
	// 服务器的名称
	Name string

	// tcp4 or other
	IPVersion string

	// 服务器绑定的IP地址
	IP string

	// 服务器绑定的端口号
	Port int
}

//============== 定义当前客户端链接的handle api ===========
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显业务
	fmt.Println("[Conn Handle] CallBackToClient ... ")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallBackToClient error ")
	}
	return nil
}

//============== 实现 ziface.IServer 里的全部接口方法 ========
// 开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	// 开启一个 go 去做服务端 listener 业务
	go func() {
		// 1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}
		// 2 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "err:", err)
			return
		}

		// 已经监听成功
		fmt.Println("start Zinx server  ", s.Name, " succ, now listenning...")

		//TODO server.go 应该有一个自动生成ID的方法
		var cid uint32
		cid = 0

		// 3 启动 Server 网络连接
		for {
			// 3.1 阻塞等待客户端建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err: ", err)
				continue
			}

			// 3.2 TODO Server.Start() 设置服务器最大连接控制，如果超过最大连接，那么则关闭此新的连接

			// 3.3 TODO Server.Start() 处理新连接的请求的 业务 方法，此时应该有 handler 和 conn 是绑定的
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++

			go dealConn.Start()

			// 我们这里暂时做一个最大512字节的回显服务
			//go func() {
			//	// 不断的循环从客户端获取数据
			//	for {
			//		buf := make([]byte, 512)
			//		cnt, err := conn.Read(buf)
			//		if err != nil {
			//			fmt.Println("recv buf err : ", err)
			//			continue
			//		}
			//		// 回显到客户端
			//		if _, err := conn.Write(buf[:cnt]); err != nil {
			//			fmt.Println("write back buf err ", err)
			//			continue
			//		}
			//	}
			//}()
		}
	}()

	//

}

// 停止网络服务
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)

	//TODO  Server.Stop() 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
}

// 开启服务器方法
func (s *Server) Server() {
	s.Start()
	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listener 的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}
}

/**
创建一个服务器句柄
*/
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}
	return s
}
