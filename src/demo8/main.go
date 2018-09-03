// demo8 project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	//  验证[]byte and string  and  which is bit byte
	var msg [8]byte
	msg[0] = 8 // ASIIC  退格
	msg[1] = 0 //  空字符
	msg[2] = 0
	msg[3] = 13
	msg[4] = 0
	msg[5] = 37 //%
	fmt.Print(msg)
	// string changeto  []byte
	s := "hello"
	fmt.Print(s, "\n")

	//  []byte change to string
	m := string(msg[:])
	fmt.Print(m)
	fmt.Print("\n")

	// string to []byte  -> string -> ASIIC code -> []byte
	a := "a"
	A := []byte(a)
	b := "b"
	B := []byte(b)
	c := "c"
	C := []byte(c)
	fmt.Print(A, "\n")
	fmt.Print(B, "\n")
	fmt.Print(C, "\n")

	aa := int(msg[5] << 8)
	//  <<  >> 这些都是位移操作。也就是把目前都二进制位数统一都左移或右移。
	fmt.Print(aa)
	fmt.Print("\n", uint16(255))

	//另外这里有一些在golang当中定义变量的方式。就是你定义的数字变量类型，就限制住了你最终能够赋上多大的值。
	var MaxInt uint16 = 1<<16 - 1
	fmt.Println("\n---", MaxInt)
	/*
		---> output:  65535
		var MaxInt uint16 = 1<<17 - 1
		fmt.Println("\n---", MaxInt)
		---> output: error 121071 overflows uint16
	*/
	var maxint int8
	maxint = 1111111111111
	fmt.Print(maxint)

}
