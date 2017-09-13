package main

import (
	"net"
	"os"
	"io"
	"bytes"
	"fmt"
)

func main(){
	if len(os.Args) != 2{
		fmt.Print("Usage: host:port")
		os.Exit(1)
	}

	service := os.Args[1]
	conn,err := net.Dial("tcp",service)
	checkError(err)
	_,err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	
	result,err := readFully(conn)
	checkError(err)
	
	fmt.Println(string(result))

	os.Exit(0)	
}

func checkError(err error){
	if err != nil {
		fmt.Print(os.Stderr,"Fatal error:%s")
		os.Exit(1)
	}
}	

func readFully(conn net.Conn)([]byte ,error){
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		_,err := conn.Read(buf[0:])
		result.Write(buf[0:])
		if err != nil {
			if err == io.EOF{	
				break
			}
			return nil,err
		}
	}
	return result.Bytes(),nil
}
