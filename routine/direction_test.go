package routine

import (
	"fmt"
	"testing"
)

// 单向channel

// ⚠️
// Go语言的类型系统提供了单方向的channel类型，分别用于只发送或只接收的channel。
// 类型chan<- int表示一个只发送int的channel，只能发送不能接收。
// 类型<-chan int表示一个只接收int的channel，只能接收不能发送。
//（箭头<-和关键字chan的相对位置表明了channel的方向。）这种限制将在编译期检测

// 因为close操作只用于断言不再向channel发送新的数据，
// 所以只有在发送者所在的goroutine才会调用close函数，
// ⚠️因此对一个只接收的channel调用close将是一个编译错误
func TestDirect(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

// 调用counter（naturals）时，naturals的类型将隐式地从chan int转换成chan<- int
// 任何双向channel向单向channel变量的赋值操作都将导致该隐式转换。
// 这里并没有反向转换的语法：也就是不能将一个类似chan<- int类型的单向型的channel转换为chan int类型的双向型的channel
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
