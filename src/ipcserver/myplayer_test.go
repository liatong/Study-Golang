package ipcserver

import (
	"fmt"
	"testing"
	"time"
)

/////  验证为什么，有一些函数一定要通过返回指针的形式来传递了。
type Student struct {
	Name string
}

func (s Student) Say() {
	fmt.Print("\nmy name is:", s.Name)
}
func (s *Student) ChangeName(name string) {
	s.Name = name
	fmt.Print("\nmy name change to :", s.Name)
}

type ClassRoom struct {
	Students map[string]*Student
}

func (cr *ClassRoom) NewStudent(id, name string) Student {
	s := Student{name}
	cr.Students[id] = &s
	return s
}

func NewClassRoom() *ClassRoom {

	//	var ssmap map[string]*Student
	ssmap := make(map[string]*Student)
	return &ClassRoom{ssmap}
}

////////
type SayHi struct {
	Name string
}

func (s SayHi) Say() {
	fmt.Print(s.Name)
}
func NewSayHi(name string) SayHi {
	ss := SayHi{name + "\n"}
	return ss
}
func (s SayHi) ChangeName(name string) {
	s.Name = name
	s.Say()
}

func TestCCenter(t *testing.T) {

	cr := NewClassRoom()
	fmt.Print(cr)
	st := cr.NewStudent("001", "li")
	cr.Students["001"].Say()
	st.Say()
	cr.Students["001"].ChangeName("changeName")
	fmt.Print("\n1: cr.students.name   2:st.say()  \n ")
	cr.Students["001"].Say()
	st.Say()

	fmt.Print("\n Testing return is not potin from a function\n")

	s := SayHi{"li"}
	s.Say()
	s1 := NewSayHi("li2")
	s1.Say()
	s1.ChangeName("li3")
	fmt.Print("debug change object: " + s1.Name + "\n")
	s1.Say()

	//NewCenterServer
	ccserver := NewCenterServer()

	//NewCClient
	cclient := CClient{ccserver}
	fmt.Print(cclient)
	cclient.LoginUser("liwentong")
	cclient.LoginUser("username001")
	//	cclient.LoginUser("username002")
	//	cclient.LoginUser("username003")
	//	cclient.LoginUser("username004")
	cclient.LoginUser("username005")
	cclient.ListUser()
	cclient.LogoutUser("username005")
	cclient.ListUser()
	cclient.SendMessage("liwentong", "testmsg")

	time.Sleep(10000 * time.Microsecond)

}
