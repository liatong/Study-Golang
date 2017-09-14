package ipc

import (
	"testing"
)

type EchoServer struct {

}

func (server *EchoServer)Handle(method,params string) *Response {
	return &Response{"OK","ECHO:"+method+"~"+params}
}

func (server *EchoServer)Name()string{
	return "EchoServer"
}

func TestIpc(t *testing.T){
	server := NewIpcserver(&EchoServer{})
	
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	
	resp1,_ := client1.Call("foo","From client1")
	resp2,_ := client2.Call("foo","From client2")
	
	if resp1.Body != "ECHO:foo~From client1" || resp2.Body != "ECHO:foo~From client2" {
	     t.Error("ipcclient.call fail. resp1",resp1,"resp2",resp2)
	}
        client1.Close()
	client2.Close()
	
}

