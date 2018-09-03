package ipcserver

import (
	"fmt"
)

//这里面涉及到要使用别的包内容。也算是正式使用一下。之后再专门罗列一个目录结构来实现具体的工程结构和自身包的导入。

//基于下方内容的阐述。我们就需要一个实例player.
type Player struct {
	Name             string "name"
	HandlerIPCServer *Ipcserver
}

func NewPlayer(name string) *Player {
	phserver := &PlayerHandlerServer{name, false}
	hipc := NewIpcserver(phserver)
	return &Player{name, hipc}
	/*
		phserver := &PlayerHandlerServer{name, false}
		ipc := NewIpcserver(phserver)
		return &ipc
	*/

}

//还需要定义好远端接收到消息的时候，需要怎么处理的Server。
type PlayerHandlerServer struct {
	Name   string
	Online bool
}

func (handler *PlayerHandlerServer) Handler(method, params string) *Response {
	//swith method to do something in case
	//	fmt.Print("\n[DEBUG]: Call method", method)
	switch method {
	case "Login":
		handler.Login(params)
	case "Logout":
		fmt.Print("call method logout")
		handler.Logout()
	case "Message":
		//		fmt.Print("[DEBUG]: get message\n")
		handler.GetMessage(params)
	default:
		fmt.Println("%s User doing anything.", handler.Name)
	}
	resp := &Response{"200", handler.Name}
	resp.Body = handler.Name
	resp.Header = "200"
	return resp

}

//定义一些具体的动作细节，最终被Handler方法统一调用。
//具体需要用户具备哪些动作。
func (handler *PlayerHandlerServer) Login(name string) {
	handler.Name = name
	handler.Online = true
	fmt.Print("\nUser logint.", handler.Name)
}
func (handler *PlayerHandlerServer) GetMessage(msg string) {
	fmt.Print("\n\t[Message]:" + handler.Name + "  user Get msg:" + msg + "\n")
}
func (handler *PlayerHandlerServer) Logout() {
	fmt.Print("I'm logout.", handler.Name)
}

//************   临时的想法  ********
//本实例的方式，是我们把客户端当成是一个被动方。然后通过CentralClient来来帮忙客户端做操作。然后让CentralServer从client接收到调用。然后去判断应该怎么做？最终再把消息。
//通过channel的方式发送回到Player身上（ipcserver）身上，然后让它解析自己应该去执行什么内容。

//那是不是可以把这个被动放调换过来。CClient，客户端，新客户端都保留？channel，让他们都goroutine去。然后需要针对某个客户做什么操作，还需要切换到这个客户的goroutine上面？
//这个可能就比较不现实了。通过CClient通过channel的方式去给CClient发消息，然后解析，然后再让Player-goroutine-ipcclient  去给ipcserver发消息，然后再让它去处理？

//如果考虑到server端自主的灵活性，以及模拟和实际项目的可能方式。貌似这种情况反倒儿是一个中更好的方式。因为这种我们想改变一些服务端的业务处理，就可以自行修改了。
//但好像又说回来，客户端需要处理的东西比较单一，可以相对固定。而现在这种方式，也可以再服务端里面做一些相对比价复杂的业务逻辑。
//只是新方式，可能真的和实际项目的方式比较接近。
//**********************************

//按照我个人的理解就是。
//有一个中心服务器，用来处理事务。并且是最好能够结合用到上面的其他包已有内容。
//

func NewCenterServer() *CenterServer {
	umap := make(map[string]*Client)
	return &CenterServer{umap}
}

type CenterServer struct {
	//list client channel.本身需要记录的就是一个play。因为我们的player里面有IPCClient，所以这里也可以变成是记录IPCCLient。
	UserMap map[string]*Client
}

//有login方法
func (cs *CenterServer) LoginUser(player *Player) bool {
	//登陆的时候，创建一个player，里面涵盖一个ipcserver端，预先定义好server中的handler方法。
	//然后放到goroutine中去执行。loigng的时候 cs会创建并记录这些player的信息。当有消息过来的时候，就遍历这个列表，
	//然后通过消息通道发送这些消息给对应的player（ipcserver) ,然后ipcserver当中已经传递了Server进去，并且已经定义好了Handler是什么了，它就可以自己去处理了。
	//step 造出一个IPCClient,然后让这个client去链接ipcserver。

	//造出来一个client，然后和这个player当中的server相连，然后给它传输东西系就可以了。
	userIPCClient := &Client{}
	//	cc := player.HandlerIPCServer.Be_client_connect() 因为CC作为IPCClient，已经有连接的功能了。
	userIPCClient.Connect(player.HandlerIPCServer)
	name := player.Name
	cs.UserMap[name] = userIPCClient
	userIPCClient.Call("Login", name)
	//	go func(player *Player){
	//		让这个Player再goroutine中去实际跑起来。然后接收处理消息。
	//	}(player)
	fmt.Print("Server: Have user login,user name:", name)
	return true
}

//有登出
func (cs *CenterServer) LogoutUser(name string) bool {
	//
	//get Client from map and  Client.Call(logout,params)
	user, ok := cs.UserMap[name]
	if ok {
		user.Call("Logout", name)
		user.Close()
		delete(cs.UserMap, name)
		fmt.Print("Server: Have a user logout.")
		return true
	} else {
		fmt.Print("Server: This user is not online")
		return false
	}
}

//有发送消息
func (cs *CenterServer) SendMessage(name, msg string) bool {
	//for the  user client, and send the message to all
	u := cs.UserMap[name]
	if u != nil {
		for n, uclient := range cs.UserMap {
			//			uclient.Conn <- msg
			fmt.Print("\nServer: send message to ", n)
			uclient.Call("Message", msg)
			//			uclient.Call("Login", msg)
			//			uclient.Call("Message", msg)
		}
		return true
	} else {
		fmt.Print("\nServer: user no online can't send msg")
		return false
	}

}

//有列出所有在线用户
func (cs *CenterServer) ListUser() {
	//for and print the map

	for name, _ := range cs.UserMap {
		fmt.Println("\nonline user:", name)
	}
}

//======这里就是一个能够交互的实验界面。
type CClient struct {
	CCServer *CenterServer
}

//用于登陆用户
func (cc *CClient) LoginUser(name string) {
	p := NewPlayer(name)
	cc.CCServer.LoginUser(p)

}

//用于登出用户
func (cc *CClient) LogoutUser(name string) {
	cc.CCServer.LogoutUser(name)
}

//用于发送消息
func (cc *CClient) SendMessage(name, msg string) {
	fmt.Print("\nUser send message to server,\t" + name + "\n")
	cc.CCServer.SendMessage(name, msg)
}

func (cc *CClient) ListUser() {
	cc.CCServer.ListUser()
}

//然后再独立一个main来进行接收用户从命令行输入的命令。
//func main(){
//	//从命令行当中接收消息。

//}
