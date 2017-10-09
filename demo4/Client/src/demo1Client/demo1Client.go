package main

import (
	"fmt"
	"flag"
	"os"
	"net/rpc"
	"bufio"
	"io"
	"os/exec"
	"io/ioutil"
	"regexp"
)

var Usage = func(){
	fmt.Println("Usage: -c command -i serveraddr -p serverport")
}

type Request struct{
        Action string
}

type Response struct{
        Header  string
        Body    string
}

type Node struct{
	Name string `json:"name"`
	Zone string `json:"zone"`
	Lan  string `json:"lan"`
	IP   string `json:"ip"`
}
type Netstatus struct {
	Ping string
        Loss string
        Min  string
	Max  string
	Avg  string
}
func readCommand(infile string)(values []string,err error){
	//open file read command
	infile1 := "/root/command"
	file,err := os.Open(infile1)
	if err != nil {
		fmt.Println("open file fail!")
		return
	} 
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]string,0)
	for {
		line,isPrefix,err1 := br.ReadLine()
		if err1 != nil{
			if err1 != io.EOF {
				err = err1	
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line")
			return		
		}
		str := string(line)
		//fmt.Print(str)
		values = append(values,str)
	}
	return
}

func callGetNode(name string){
	var req Request
	var resp Node
	client,err := rpc.DialHTTP("tcp","127.0.0.1:1224")
	if err != nil{
		fmt.Println("Can't not connect to sever")
		return 
	}
	req.Action = name
	err = client.Call("Watcher.GetNode",req,&resp)
	if err != nil {
		fmt.Println("Call server function error")
		return 
	}	
	fmt.Println(resp)
}
func callCommand(command string,c chan Response){
	var req Request
	var resp Response
	//call server Getinfo
	client,err := rpc.DialHTTP("tcp","127.0.0.1:1224")
	if err != nil{
		fmt.Println("Can't not connect to sever")
		return 
	}
	//call the function. need args server obj function,request,response	
	//req := Request{"This is action"}
	req.Action = command
	//fmt.Print("----client call server----\n")
	err = client.Call("Watcher.Getinfo",req,&resp)
	if err != nil {
		fmt.Println("Call server function error")
		return 
	}	
	//fmt.Println(resp,"\n")
	//c <- "200"
	c <- resp
}

func detectNet(ipaddr string)map[string]string{
     cmd := "ping -q -c 10" + ipaddr
     fmt.Println("cmd is :",cmd)
     m := make(map[string]string)
     m["1"] = "abcd"
     return  m
}

func main(){
	command := flag.String("c","getinfo","Get server info")
	server_addr := flag.String("i","127.0.0.1","Server ip addr")
	server_port := flag.String("p","1224","Server port")
        flag.Parse()
	
	//check the args num
	args := os.Args[1:]
	if args == nil || len(args) < 6 {
		Usage()
		return
	}
	fmt.Printf("Client command is: %s\n",*command)
	fmt.Printf("Server ip addr is: %s\n",*server_addr)
	fmt.Printf("Server listen port is: %s\n",*server_port)	

	// define a response struct. only use int struct at there

	//get commands
	commands,_ := readCommand("/root/command")
	fmt.Print(commands)
	c := make(chan Response)

	for _,command := range commands{
		fmt.Print("-------")
		go callCommand(command,c)
                fmt.Print(command,"\n")
	}
	for i:=0;i<len(commands);i++{
	   select {
	     case resp := <- c:
		 fmt.Print(resp)
	         fmt.Print("--Have a gorouting finesh--\n") 
	     //default:
 	     //    fmt.Print("This is a default")
	   }
	}
	fmt.Print("----All Call command is end -----\n")
        //get Node
        //nodes,_ := readCommand("/root/node")
	//fmt.Print(nodes)
	//for _,node := range nodes{
	//	go callGetNode(node)
        //        fmt.Print(node,"\n")
	//}
	
	//------Here need to wait all go return---------//

	//------Run os command---------//

	cmd := exec.Command("/bin/bash","-c","ping -q -c 5 192.168.1.1")
	
	stdout,err := cmd.StdoutPipe()
	if err != nil {
	    fmt.Println("can not obtain cmd stdout")
	    os.Exit(1)	
	}
	
	if err:= cmd.Start();err != nil{
	    fmt.Println("Error: The command is err",err)
	    os.Exit(1)	
	}
	
	bytes,err := ioutil.ReadAll(stdout)
	if err != nil {
	    fmt.Println("readall stdout",err.Error())
	    os.Exit(1)	
	}
	if err := cmd.Wait();err != nil {
	    fmt.Println("wait")
	    os.Exit(1)	
	}
	fmt.Println("Command stdout :%s",bytes)
	fmt.Println("Command stdout:",string(bytes))
	
	regmsg := "errors, 100% packet loss"
	reg := regexp.MustCompile(`(\w+)% packet loss`)
	reg2 := regexp.MustCompile(`min/avg/max/mdev = (\w+\.\w+)/(\w+\.\w+)/(\w+\.\w+)/(\w+\.\w+)`)
	reg1 := regexp.MustCompile(`losaaas`)
        ss := reg.FindStringSubmatch(string(bytes))
	if err != nil {
	    fmt.Println("can not match regex")
	}
        ss2 := reg2.FindStringSubmatch(string(bytes))
	fmt.Println(ss)
	fmt.Println(ss2)
	fmt.Printf("%q",reg.FindStringSubmatch(regmsg))
        ss3 := reg1.FindAllString(regmsg,-1)
        fmt.Println(ss3)
        if ss3 == nil {
            fmt.Println("[] == nil")
	}
	fmt.Println(reg1.FindAllString(regmsg,-1))
	m := detectNet("test")
        fmt.Println(m)

}
