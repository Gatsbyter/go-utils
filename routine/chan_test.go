// CSP是一种现代的并发编程模型
// CSP 是 “顺序通信进程”(communicating sequential processes) 的简称
// goroutine和channel 就支持“顺序通信进程”

package routine

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	var ch1 chan int         // 定义   ch1是nil
	ch2 := make(chan int)    // 初始化 ch2不是nil 但这是个无缓冲区的channel
	ch3 := make(chan int, 0) // 初始化 和ch2一样
	ch4 := make(chan int, 5) // 初始化 分配空间 带缓冲区的channel

	fmt.Printf("%v\n", ch1) // nil
	fmt.Printf("%v\n", ch2) // not nil
	fmt.Printf("%v\n", ch3) // not nil
	fmt.Printf("%v\n", ch4) // not nil

	ChCh(ch4)

	close(ch4) // 内置函数 用来关闭channel
}

// ⚠️当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用
// 因此调用者和被调用者将引用同一个channel对象
func ChCh(ch chan int) {
	// 一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞
	// 直到另一个goroutine在相同的Channels上执行接收操作
	// 当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句。
	//
	// 反之，如果接收操作先发生，那么接收者goroutine也将阻塞，
	// 直到有另一个goroutine在相同的Channels上执行发送操作

	// 所以 这里直接死锁了
	ch <- 2   // 将数据发送到channel
	x := <-ch // 定义一个变量接受channel数据
	<-ch      // 从channel取出数据 但是数据被抛弃

	fmt.Println(x)
}

// 基于无缓存Channels的发送和接收操作将导致两个goroutine做一次同步操作。
// 因为这个原因，无缓存Channels有时候也被称为同步Channels。
// 当通过一个无缓存Channels发送数据时，接收者收到数据发生在唤醒发送者goroutine之前
