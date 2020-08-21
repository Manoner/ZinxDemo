package ziface

/*
   IRequest 接口：
   实际上是把客户端 请求的连接信息 和 请求的数据 包装到了 Request里
*/

type IRequest interface {
	GetConnection() IConnection // 获取请求的连接信息
	GetData() []byte            // 获取请求信息的数据
}
