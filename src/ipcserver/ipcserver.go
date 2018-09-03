// ipcserver project ipcserver.go
package ipcserver

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Header string "header"
	Body   string "body"
}

type Server interface {
	Handler(method, paras string) *Response
}

type Ipcserver struct {
	Server
}

func NewIpcserver(s Server) *Ipcserver {
	ipc := &Ipcserver{s}
	return ipc
}

//理解服务器就是应该listen，然后有链接了，就触发一个协程去处理具体的处理。只是这里我们的client.connect()没办法真的通过网络链接上。只好是在client.connect()的这个方法
//里面，再传递一个server本身的chan给它,并且这时候，因为client不是真的在网络发起链接请求，所以server的listen没办法真正的触发。所以这时候，listen就得有机制假装
//接收到了请求，然后再去创建出一个协程去处理。
func (ipc *Ipcserver) listen(client chan string) {
	//不断循环内部存储中但chan是否有数据写入，如果有
	go func(client chan string) {
		var req Request
		for {
			msg := <-client
			if msg == "CLOSE" {
				close(client)
				break
			}
			//就处理具体的事务。包括底层的字符串解析.这里先是将接收到的byte数组给解析出来。然后再做判断。
			err := json.Unmarshal([]byte(msg), &req)
			if err != nil {
				fmt.Print("Invalid request format:", msg)
			}
			//
			resp := ipc.Server.Handler(req.Method, req.Params)
			b, err := json.Marshal(resp)
			client <- string(b)

		}
		fmt.Print("session is close")
	}(client)

}
func (ipc *Ipcserver) Be_client_connect() chan string {
	cc := make(chan string)
	ipc.listen(cc)
	return cc
}

// client 本身就应该要具备有这个连接存在，当需要call的时候，就把消息发送出去。而我们这里也仅仅是把这个替换成为channel而已。
type Client struct {
	Conn chan string
}

func (c *Client) Connect(server *Ipcserver) {
	c.Conn = server.Be_client_connect()
}
func (c *Client) Call(method, paras string) (resp *Response, err error) {
	//就是把这个method和params封装好，然后通过通道传输出去，然后就等待响应回传回来了。其中涉及到底层的封装。
	req := &Request{method, paras}
	//把这个req转换成为json，然后通过c.con传递出去。
	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		return
	}
	c.Conn <- string(b)
	r := <-c.Conn
	var respl Response
	//将返回来的byte数组装成json字符串，并且装载到我们的json类型对象中去。
	err = json.Unmarshal([]byte(r), &respl)
	resp = &respl
	return
}

func (c *Client) Close() {
	c.Conn <- "CLOSE"
}

//  小结一下： 对于一个未知的不知道如何的程度，确实应该先分析分析业务，然后分析分析用例，然后才是从中得到实体。
//	然后就是从这些实体当中抽象出，能够抽象的对象。
//	然后就是针对这些对象去理清楚，彼此之间的关系，以及对象本身应该具备什么属性，应该具备什么方法。什么样的沟通方式才会是最佳的。
//	然后等到，你基本的熟悉这种做法之后，积累出来的拆解方法，就会慢慢成为你的一种本能反应。或习惯性思维了。
//	最终你就能够很容易面对业务，拆解，落地实现了。
