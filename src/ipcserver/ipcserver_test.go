package ipcserver

import (
	"fmt"
	"testing"
	"time"
)

//这里面真的如果自己没有认真的去写过，是很难发现里面的问题。
//写一个handler，只要实现了Server接口，我们就能够当成是一个Server被使用。
//这样就能够造出一个ipcserver了。

//然后下一步就是再造出一个client出来。然后让client去connect

type EchoServer struct {
}

func (e *EchoServer) Handler(method, params string) *Response {
	return &Response{"200\n", method + "  " + params}

}

func TestIcp(t *testing.T) {
	echoserver := &EchoServer{}
	echoicp := NewIpcserver(echoserver)

	//造出来了，EchoServer 之后，我们就还要造一个client出来。
	id := 0
	for {
		go func(echoicp *Ipcserver, id int) {
			c1 := Client{}
			c1.Connect(echoicp)
			resp, _ := c1.Call("   hello", "world")
			fmt.Print(id)
			fmt.Println("\n" + resp.Header + resp.Body)
			//			fmt.Println("------\n")
		}(echoicp, id)
		id++
		if id > 3 {
			break
		}
	}
	time.Sleep(3000 * time.Millisecond)
}
