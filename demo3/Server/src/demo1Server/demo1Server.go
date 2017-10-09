package main 

import (
	"fmt"
	"io"
	"os"
	"net"
	"net/rpc"
	"net/http"
	"io/ioutil"
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
        logfile := "/tmp/testlogfile"
        b,err := ioutil.ReadFile(logfile)
	if err != nil {
		panic(err)
	}
        conf := string(b)
	(* resp).Header = "Header:header"
	(* resp).Body = "Body:command is"+req.Action+"\nConfig file"+conf
	return nil
}
func checkFileExist(filename string)bool{
	if _,err := os.Stat(filename);os.IsNotExist(err){
		return false
	}
	return true
}
func main(){
	watcher := new(Watcher)
	rpc.Register(watcher)
	rpc.HandleHTTP()
        
	logFile := "/tmp/testlogfile"
	//----check file exist and create file ----//
 	if exist := checkFileExist(logFile);exist{
		fmt.Println("Log file is exist")	
	}else{
		logfile,err := os.Create(logFile)
		defer logfile.Close()
		if err != nil{
		    fmt.Println("Create logfile error")
                    os.Exit(1)
		}
		fmt.Println("Create logfile success!")
	}
	//------Read file use file.Read(buf)-------//
        f,err := os.Open(logFile)
	defer f.Close()
	if err != nil {
		fmt.Println("file cant open")
	}
	buf := make([]byte,1024)
        for {
	    n,err := f.Read(buf)
	    if err != nil && err != io.EOF{
	       panic(err) 
	    }
 	    if n == 0 {
	       break
	    }
	    os.Stdout.Write(buf[:n]) 
	}
	fmt.Println("========")

	//-------Read file use ioutil---------//
        d,err := os.Open(logFile)
	defer d.Close()
        b,err := ioutil.ReadAll(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

    
	l,err := net.Listen("tcp",":1224")
	if err != nil {
		fmt.Println("Liston fail")
		return
	}
	fmt.Println("Server listen at 1224")
	http.Serve(l,nil)	

}
