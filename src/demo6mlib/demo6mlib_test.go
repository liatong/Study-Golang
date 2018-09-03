package demo6mlib

import (
	"fmt"
	"testing"
)

func TestOps(t *testing.T) {

	//这里原先就是从函数里面传递
	mm := NewMusicManager()
	fmt.Print(mm, "\n")
	d := mm.Len()
	fmt.Print(d)
	m1 := &MusicEntry{"001", "oleole", "aircondition", "oleole", "MP3"}
	mm.Add(m1)
	//https://gocn.vip/question/1519  这里面可以看看别人如何解释&取地址，* p  取指针p所对应的值。 mm是个指针，取这个指针对应的具体值。
	//这里还是要记住，对于map这种对象来说，要传递都是要传递指针。
	//&mm  这个是取指针都意识。
	//这样的形式就是不行。 从newmm开始，需要的是一个指针类型。所以mm接收到的是一个指针。mm本身是个变量。mm. 这种方式本身会引用。但如果要把mm再传递给别人（也就是说newmm
	//里面早出来的那个musicmanager的话，那就还是需要musicmanager已开始创造出来的地址。）现在这个地址，也就是mm这边变量里面存储的指针地址了。所以*mm 就是取变量mm的值。
	//pl := &Player{mm}
	pl := &Player{*mm}
	pl.Play("oleole")

	//如果按照上面的理解，我这里player构造的时候，不是一个指针，那表示传递的值类型。那是不是player里面的mm 和 外面的mm 已经不是同一个mm了。
	//小结：其实最终 m := &a  m就是一个指针变量，这个变量里面存放的是一个指针。而*m  针对指针变量，那么传递给其他人就是一个指针了。 p := *m 那p就也是一个指针变量。
	fmt.Print("\n", pl.musicm.Len(), mm.Len())

	fmt.Print("\nsame the len\n")
	fmt.Print("\nchange the mm len\n")
	mm.Del("oleole")
	fmt.Print("\n", pl.musicm.Len(), mm.Len())
	fmt.Print("\nalways same the len\n")
	mm.Add(m1)

	fmt.Print("\n", pl.musicm.Len(), mm.Len())
	fmt.Print("\n&mm is what?-->", &mm)
	fmt.Print("\nmm is a point var:-->", mm)
	p := *mm
	fmt.Print("\np := *mm and this p also is a point var:-->", p)
	fmt.Print("\np is not a point var. but *p is a what? --> p (type MusicManager) \n")
	fmt.Print("&p also is a point. so that &p = mm:--> ", &p)
	//&取了对象的地址，交给别人，别人是指针变量。 mm是一个指针变量，存的是指针。   *mm 地址变量被解开，变回来了具体的对象。p就又编程是一个变量而已。存储着对象的地址。

	//mm本身是个变量，有自己的地址空间。mm里面存的是别人的地址。这个地址指向的是一个对象。能引用。
	// 0x00000（mm)-> 0x30000 -> object
	//*mm 取地址指向的具体值。  p := *mm  . p = 0x30000(p)-->object

}
