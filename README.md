# ZinxDemo

记录zinx框架的学习过程

### V0.4 版本执行
实现全局参数的配置和读取
```bash
➜  ZinxDemo git:(master) ✗ go run apptest/server/server.go
Add Router success! 
[START] Server name: demo server,listenner at IP: 127.0.0.1, Port 7777 is starting
[Zinx] Version: V0.4, MaxConn: 3,  MaxPacketSize: 12000
start Zinx server   demo server  succ, now listenning...
```


### V0.5 消息封装

封包格式
![img](http://cdn.note.manoner.com/2-TCP%E7%B2%98%E5%8C%85%E9%97%AE%E9%A2%98-%E6%8B%86%E5%8C%85%E5%B0%81%E5%8C%85%E8%BF%87%E7%A8%8B.jpeg)


### V0.6 多路由模式


### V0.7 读写分离模型
![img](http://cdn.note.manoner.com/3-Zinx-V0.7%E6%9E%B6%E6%9E%84%E6%A8%A1%E5%9E%8B.jpeg)
