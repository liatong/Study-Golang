// demo5 project main.go
package main

import (
	"fmt"
)

type Interger int

func (a Interger) Less(b Interger) bool {
	return a < b
}

func main() {
	fmt.Println("Hello World!")

	//可以给任何类型加上自己的方法，因为类型都是值类型，值类型就能够加上方法。
	var a Interger = 1
	if a.Less(2) {
		fmt.Println("a less 2")
	}
	// 面向对象和面向过程其实就是一个语法糖而已。
	//面向过程： type interger int   ->  func less ( integre a,b)  return a>b  调用 less(a,b)
	//面向对象： type integre  int  fun(a integre) less(b integre) return a>b  调用 a.less(b)

	////那这里就总结都是值类型。

	//===========值语义和引用=========
	//另外一项很重要的，golang 和 c 一样，类型传递都是基于值类型的。比如字符型，布尔型。就算是map等也都是一样都。
	//如果想要改变一个变量都值，go和c一样就都需要以传递指针。在构造函数的时候，也同样需要这样的方式。
	//func (a *类型指针）method-name(a int) bool {  *a =  *a + b }
	// &a 为取a的指针。   *a 为取a指针的具体值。

	//========== 一项注意的 =======
	//如某一个静态的产品，可是一旦加入，就必须把一些事情给好了。

	c := &Car{"car"}
	c.Run()
	hc := &Housecar{*c, "housecar"}
	hc.Run()

}

//---------  结构体   ------
//在go语言里面的结构体和其他的类是一样的效果，因为可以给结构体体加上具体的方法。
//但是没有继承，只有组合。 如你需要一个别的结构体的功能，那就把它组合进来，在你自己的结构体中存放它。然后可以在基于被组合的结构体方法，在构造出属于自己的方法。

//定义一个车的结构体，能够跑。然后再定义一个马的结构体，能跑还能叫。

type Car struct {
	name string
}

func (c *Car) Run() {
	fmt.Print("i'm a car, i can run\n")
}

type Housecar struct {
	Car
	name string
}

func (hc *Housecar) Run() {
	fmt.Print("i'm a house and ")
	hc.Car.Run()
}

//-------go 语言中一个非常重要的特性就是接口。-------
//如果说类型系统是基石，那么接口就是基石的基础。

//其他语言的接口，接口主要作为不同组建之间的契约存在。对契约的实现是强制性的。
//go语言的接口，却不是强制性的需要继承。 比如接口1，写的类就只能是适合接口1，接口2写的类就只能适合接口2。
//go的接口是写了一个类，实现了一些方法。 如果这些方法在其他的接口中有实现，那么就可以进行赋值了，说明这个类是实现了此接口的。
//比如： type document struct{}   func(a *d)amethod{} | type in interface{ amethod()()}
//var in1 in = new(document)
// in1.amethod()
//print "i'm a house"

//----------------- 接口组合的方式 --------
//接口的这种灵活性，是之前很多语言所不具备的。所以这是一个要让语言编写出灵活架构的核心所在。
//需要注意的一点就是，要继承别人的类。在继承接口时，如果类本身的方法是需要传递指针的。那么继承接口的方式也得是拿指针的方式
// type Inte int  func(a Inte)less  func(a *Inte)add  func  | type LessAddr interface { less(),add()}
//如 var b LessAdder = &a  |  这样对于哪些有指针的方法才会依据有效。如果是传值的话，那原本方法中的传指针就不能够使用了。

//接口查询
//var file1 Write = ...
// if file5,ok := file1.(two.Istream);ok{   ....
//}

//----接口组合---
// 已经有了read,write接口
//readwrite 接口只需要   type readwrite interface { read  write }
//
