package ipc

import "encoding/json"

type IPCClient struct {
	conn chan string
}

func NewIpcClient(server *IPCServer) *IPCClient {
	c := server.Connect()
	return &IPCClient{c}
}

func (client *IPCClient) Call(method, params string) (resp *Response, err error) {
	req := &Request{Method:method, Params:params}
	var b []byte  // 输出一下看看究竟是什么格式？？
	b, err = json.Marshal(req)  //err没有定义这个变量啊？？err被视为是特殊的变量吗？返回值里面定义了这个变量，函数体里面可以直接用。
	if err != nil{
		return
	}
	client.conn <- string(b)
	str := <- client.conn  // 书上写的等待返回值？？ 为什么要先传给client.conn这个变量。有解释，但是现在我看不懂。
	// 怎么处理的请求，如何转化为响应的？？是因为协程的原因
	var resp1 Response  // 这个地方为什么不直接定义指针啊？定义指针就没有开辟存储的空间吗？？
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1
	return
}

func (client *IPCClient)Close() {
	client.conn <- "CLOSE"
}