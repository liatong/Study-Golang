package main 

import (
	"fmt"
//	"io"
	"net"
	"net/rpc"
	"net/http"
)
type Request struct{
	Action string	
}
type Response struct{
	Header	string
	Body	string
}

type Watcher int
func(w *Watcher)Getinfo(req Request,resp *Response)error{
	fmt.Print(req)
	(* resp).Header = "Header:header"
	(* resp).Body = "Body:command is"+req.Action
	return nil
}

func main(){
	watcher := new(Watcher)
	rpc.Register(watcher)
	rpc.HandleHTTP()

	l,err := net.Listen("tcp",":1224")
	if err != nil {
		fmt.Println("Liston fail")
		return
	}
	fmt.Println("Server listen at 1224")
	http.Serve(l,nil)	

}
