// demo7 project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//
	for i := 0; i < 10; i++ {
		go Add(i, 1)
	}
	//如果只有以上这些代码，那么程序所打印出来的数据是不全的，这些协程虽然都有创建出来。但是我们的主进程不会等这些协程都执行完成已经退出主进程了。
	//所以我们的主进程需要在这里等待我们的协程执行都完成才能够退出主进程。

	c := make(chan int)
	//定义一个chan是什么类型。
	//定义一个chan的map
	// cm = make(map[string] chan int)
	//var cm map[string] chan int
	for i := 0; i < 100; i++ {
		go AddChannel(i, 1, c)
	}
	t := 0
	for {
		<-c
		t = t + 1
		fmt.Print("time:\n", t)
		if t >= 100 {
			break
		}
	}
	//
}

func Add(x, y int) {
	z := x + y
	fmt.Print("x+=y\n", z)
}

//使用channel的方式来重写这个功能。
func AddChannel(x, y int, c chan int) {
	z := x + y
	fmt.Print("channel x+y=\n", z)
	c <- 1=
	 2145
}

// 最好能够实际写一个支撑大量并发但小程序。
