package main

import (
	"fmt"
	"flag"
	"os"
	"net/rpc"
	"bufio"
	"io"
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
func readCommand(infile string)(values []string,err error){
	//open file read command
	infile1 := "/tmp/command"
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
		fmt.Print(str)
		values = append(values,str)
	}
	return
}

func callCommand(command string){
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
	err = client.Call("Watcher.Getinfo",req,&resp)
	if err != nil {
		fmt.Println("Call server function error")
		return 
	}	
	fmt.Println(resp)
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
	commands,_ := readCommand("/tmp/command")
	fmt.Print(commands)
	for _,command := range commands{
		go callCommand(command)
                fmt.Print(command,"\n")
	}

}
