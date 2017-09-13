packge ipc 

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "Code"
	Body string "body"
}

type server interface {
	Name() string
	Handle(method,params string) *Response
}

type Ipcserver struct {
	server
}

func NewIpcserver(server server)*Ipcserver{
	return &Ipcserver{server}
}

func (server *Ipcserver)Connect()chan string{
	session := make(chan string,0)

	go func(connecting chan string){
		for {
			revice := <- connecting
			if revice == "CLOSE" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(revice),&req)
			if err != nil{
				fmt.Println("error format")
			}
			resp := server.Handler(req.Method,req.Params)
			// resp is struct need encode to byte steam let it can be send as a string to channel.
			b,err := json.Marshal(resp)
			connectiong <- b
		}
		fmt.Println("Connect close.")
	}(session)

	fmt.Println("A new session has been create successful!")
	return session

}



