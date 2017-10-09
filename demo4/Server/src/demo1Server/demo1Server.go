package main 

import (
	"fmt"
	"io"
	"os"
	"io/ioutil"
	"net"
	"net/rpc"
	"net/http"
	"encoding/json"
	"bufio"
)
type Request struct{
	Action string	
}
type Response struct{
	Header	string
	Body	string
}
type Node struct{
	Name string `json:"name"`
	Zone string `json:"zone"`
	Lan  string `json:"lan"`
	IP   string `json:"ip"`
}

type Watcher int
func(w *Watcher)Getinfo(req Request,resp *Response)error{
	//fmt.Println("Requst from cline:",req)
        configfile := "/root/config"
        config,err := ioutil.ReadFile(configfile)
        if err != nil {
            config = nil
        } 
        fmt.Println(config)
	(* resp).Header = "{'Header':{'Type':'Config'}}"
	(* resp).Body = "{'Body':"+string(config)+"}"
	return nil
}

func(w *Watcher)GetNode(req Request,resp *Node)error{
	fmt.Println("Requst from cline:",req)
        configfile := "/root/config"
        rf,err := os.Open(configfile)
	defer rf.Close()
        if err != nil {
	    return err
	} 
  	buf := bufio.NewReader(rf)
	var n Node
	for {
	    line,err := buf.ReadString('\n')
	    if err != nil {
	        if err == io.EOF {
		    fmt.Println("------")
		    break
		}
		os.Exit(1)
	    }
            //tjstring :=`{"name":"Node1","zone":"zone2","lan":"1","ip":"192.168.1.1"}`
	    fmt.Println(line)
            err = json.Unmarshal([]byte(line),&n)
            if err != nil {
	        return err
            }
            fmt.Println("Node struct n.Name is:")
            fmt.Println(n.Name)
	    fmt.Println("----line----")
	}
	return err
}
func checkFileExist(file string)bool{
    if _,err := os.Stat(file);os.IsNotExist(err){
	return false
    }
    return true
}

func main(){
	watcher := new(Watcher)
	rpc.Register(watcher)
	rpc.HandleHTTP()
        //-----check file exist and create config file-----//
        configfile := "/root/config"
        if  ! checkFileExist(configfile){
            _,err := os.Create(configfile)
	    if err != nil{
	   	fmt.Println("can't create config file")
		os.Exit(1)
	    }
	}
	f,err := os.Open(configfile)
	defer f.Close()
        if err != nil{
            fmt.Println("config can't read")
	    os.Exit(1)
	}

	//------Read config file------//
        b,err := ioutil.ReadFile(configfile)       
        if err != nil {
            panic(err)
        }
        fmt.Println(string(b)) 

        //-----Read pea line from file---//
        rf,err := os.Open(configfile)
	defer rf.Close()
        if err != nil {
	    os.Exit(1)
	} 
  	buf := bufio.NewReader(rf)
	for {
	    line,err := buf.ReadString('\n')
	    if err != nil {
	        if err == io.EOF {
		    fmt.Println("------")
		    break
		}
		os.Exit(1)
	    }
	    fmt.Println("----line----")
	    fmt.Println(line)
	}
         
        //----from struct to byte[]  to json string ----//
        var n Node
        n1 := Node{"Node2","zone1","1","192.168.1.2"}
        bn1,err := json.Marshal(n1)
        fmt.Println(bn1)
        jn1 := string(bn1)
        fmt.Println("------------")
        fmt.Println(jn1)
 
        //-----from json string to struct ---// 
        tjstring :=`{"name":"Node1","zone":"zone2","lan":"1","ip":"192.168.1.1"}`
        err = json.Unmarshal([]byte(tjstring),&n)
        if err != nil {
          fmt.Println("Can'g direct change")
        }
        fmt.Println("Node struct n.Name is:")
        fmt.Println(n.Name)

	l,err := net.Listen("tcp",":1224")
	if err != nil {
		fmt.Println("Liston fail")
		return
	}
	fmt.Println("Server listen at 1224")
	http.Serve(l,nil)	

}
