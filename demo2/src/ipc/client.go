package ipc

import (
	"encoding/json"
)

type Ipcclient struct {
	conn chan string
}

func NewIpClient(server *Ipcserver) *Ipcclient{
      s := server.Connect()
      return &Ipcclient{s}
}

func (client *Ipcclient)Call(method,params string)(resp *response,err error){
	req := &Requst{method,params}
	var b []byte
	b,err = json.Marshal(req)
	if err != nil{
		return 
	}
	client.conn <- string(b)
	str <- client.conn
	var resp1 Response
	err = json.Unmarshal([]bypte(str),&resp1)
	resp = &resp1
	return 
}

func (client *Ipcclient)Close(){
	client.conn <- "CLOSE"
}