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
	"time"
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
	//infile1 := "./command"
	values = make([]string,0)
	file,err := os.Open(infile)
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
		fmt.Printf("Test:GetNode:readCommand----%s",str)
		//fmt.Printf("Test:GetNode:readCommand len is ----%d",len(str))
		values = append(values,str)
	}
	return
        //------another read pear line from file method ----//
	//infile1 := "./command"
	//values = make([]string,0)
        //rf,err := os.Open(infile)
	//defer rf.Close()
        //if err != nil {
	//    return nil,err
	//} 
  	//buf := bufio.NewReader(rf)
	//for {
	//    line,err := buf.ReadString('\n')
	//    if err != nil {
	//        if err == io.EOF {
	//	    fmt.Println("------")
	//	    break
	//	}
	//        return nil,err
	//    }
	//    str := string(line)
	//    fmt.Printf("Test:GetNode:readCommand----%s",str)
	//    fmt.Printf("Test:GetNode:readCommand len is----%d",len(str)
	//    values = append(values,str)
	//}
        //return values,nil 
}

func callGetNode(name string,cn chan Node){
	var req Request
	var resp Node
	client,err := rpc.DialHTTP("tcp","127.0.0.1:1224")
	if err != nil{
		fmt.Println("Can't not connect to sever")
		return 
	}
        fmt.Printf("Test:GetNode:NodeName ---- name is %s\n",name)
	req.Action = name
        fmt.Printf("Test:GetNode:NodeName ---- name len %d",len(name))
        fmt.Printf("Test:GetNode:NodeName ---- name len %d",len("Node1"))
	err = client.Call("Watcher.GetNode",req,&resp)
	if err != nil {
		fmt.Println("Call server function error")
		return 
	}	
        fmt.Printf("Test:GetNode:Response---Response Node is: %s\n",resp.Name)
        cn <- resp
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

func realywork(u string ,queue chan string,backchannel chan string){
      fmt.Printf("Realywork--------func realywork be exec:%s------\n",u)
      time.Sleep(2*time.Second)
      fmt.Printf("Realywork-------func  realywork end delete one from queue------\n")
      <- queue
      backchannel <- u
      time.Sleep(3*time.Second)
      
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

	//-------get commands--------//
	//commands,_ := readCommand("./command")
	//fmt.Print(commands)
        //// -- test make len(vars) --
        var urls = []string{"01","02","03","04","05","06","07","08","09","10","11","12","13","14","15","16","17","18","19"} 
	//c := make(chan Response,len(urls))

	//for _,command := range commands{
	//	fmt.Print("-------")
	//	go callCommand(command,c)
        //        fmt.Print(command,"\n")
	//}
	//for i:=0;i<len(commands);i++{
	//   select {
	//     case resp := <- c:
	//	 fmt.Print(resp)
	//         fmt.Print("--Have a gorouting finesh--\n") 
	//     //default:
 	//     //    fmt.Print("This is a default")
	//   }
	//}
	//fmt.Print("----All Call command is end -----\n")

        //-------- test get Node --------//
        nodes,_ := readCommand("./node")
        cn := make(chan Node)
	fmt.Printf("Test:GetNode:Read all node list is:%s",nodes)
	for _,node := range nodes{
                fmt.Printf("Test:GetNode:Node ---node name:%s",node) 
                fmt.Printf("Test:GetNode:Node len ---node name:%d",len(node)) 
		go callGetNode(node,cn)
                fmt.Print(node,"\n")
	}

	for i:=0;i<len(nodes);i++{
	   select {
	     case resp := <- cn:
		 fmt.Print(resp)
	         fmt.Print("Test:GetNode:--Have a gorouting finesh--\n") 
	     //default:
 	     //    fmt.Print("This is a default")
	   }
	}
	
	//------Here need to wait all go return---------//

	//------Run os command---------//

	fmt.Print("----Testing exec.command ping -----\n")

	cmd := exec.Command("/bin/bash","-c","ping -q -c 5 127.0.0.1")
	
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
	fmt.Print("----End testing exec.Command ping -----\n")

	//--------Test goroutime queue ------//
       //var urls = []string{"01","02","03","04","05","06","07","08","09","10","11","12","13","14","15","16","17","18","19"} 
       backchannel := make(chan string,len(urls))
       queue := make(chan string,5)

       go func(){
            for _,u := range urls {
		 fmt.Printf("Queue------start to  go func to run:%d--------\n",len(queue))	
	 	 queue <- u
		 fmt.Printf("Queue------Queue not full can  add go func to run --------\n")	
		 go  realywork(u,queue,backchannel )
            }
       }()
       for i:=0;i<len(urls);i++{
          select {
            case resp := <- backchannel:
                fmt.Printf("Main--%s:Have a gorouting finesh--\n",resp) 
            //default:
            //    fmt.Print("This is a default")
          }
       }

}
