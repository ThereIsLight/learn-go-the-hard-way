package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string `json:"method"`
	Params string `json:"method"` //使用golang的struct tag进行json和db映射的时候??
}
// 在golang中叫标签（Tag），在转换成其它数据格式的时候，会使用其中特定的字段作为键值。
// 将method、params改为小写，则导出的接送字符串中没有这两个键值。相见learnjson.go

type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

type Server interface {
	Name() string
	Handle(method, params string) *Response  // 接受的参数都是request的成员变量
}

type IPCServer struct {
	Server
}

func NewIpcServer(server Server) *IPCServer {
	return &IPCServer{server}
}  // ??这个地方实在是看不懂了？？ 接口类型到底是什么？？

func (server *IPCServer)Connect() chan string {  // return chan??
	session := make(chan string, 0)  //
	// 定义并且调用匿名函数 c为函数定义的形式参数，session为函数调用时传入的实际参数。
	// 协程没有返回值
	go func(c chan string) {
		for {
			request := <- c
			if request == "CLOSE" {  //client端传过来的就是CLOSE
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			//1.Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构。
			//第二个参数必须是指针，否则无法接收解析的数据。但是这个地方只是对变量取地址罢了，不是整正的指针啊。
			if err != nil {
				fmt.Println("Invaild request format:", request)
				return
			}
			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)  // 指针变量，编码时自动转换为它所指向的值。只不过指针更快，且能节省内存空间。
			c <- string(b)
		}
		fmt.Println("Session closed")
	}(session)
	fmt.Println("A new Session has been created successfully.")
	return session
}