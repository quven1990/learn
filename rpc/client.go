package main

import (
	"fmt"
	//"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//用rpc链接服务器
	//conn, err := rpc.Dial("tcp", "127.0.0.1:8088")
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()
	var resp *string
	conn.Call("hello.HelloWorld", "小饭", &resp)
	fmt.Println(*resp)
}
