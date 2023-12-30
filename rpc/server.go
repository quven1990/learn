package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Hello struct {
}

func (this *Hello) HelloWorld(name string, resp *string) (err error) {
	*resp = name + "您好！"
	return err
}

func main() {
	//注册rpc服务，指定对象和方法
	err := rpc.RegisterName("hello", new(Hello))
	if err != nil {
		fmt.Println("注册 err:", err)
		return
	}
	//设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listener.Close()
	//建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept err:", err)
		return
	}
	defer conn.Close()
	//绑定服务
	//rpc.ServeConn(conn)
	jsonrpc.ServeConn(conn)
}
