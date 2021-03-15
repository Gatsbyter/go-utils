package routine

import (
	"fmt"
	"testing"
	"time"
)

// select语句switch语句稍微有点相似，也会有几个case和最后的default选择支。
// 每一个case代表一个通信操作(在某个channel上进行发送或者接收)并且会包含一些语句组成的一个语句块。
//
// select会等待case中有能够执行的case时去执行。
// 当条件满足时，select才会去通信并执行case之后的语句；这时候其它通信是不会执行的。
// 一个没有任何case的select语句写作select{}，会永远地等待下去。
func TestSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int, 10)
	y := 10000

	//go func() {
	//	ch2 <- 3
	//	time.Sleep(time.Hour)
	//}()
	//
	//go func() {
	//	ch1 <- 3
	//	time.Sleep(time.Hour)
	//}()
	//
	//go func() {
	//	<- ch3
	//	time.Sleep(time.Hour)
	//}()

	time.Sleep(time.Second)

	// 如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会
	// 如果没有default select会阻塞直到有一个分支返回
	for {
		select {
		case <-ch1:
			fmt.Println(11111)
		case x := <-ch2:
			fmt.Println(22222, x)
		case ch3 <- y:
			fmt.Println(33333)
		default:
			fmt.Println(44444)
			return
		}
	}
}

// ch这个channel的buffer大小是1，所以会交替的为空或为满，
// 所以只有一个case可以进行下去，无论i是奇数或者偶数，它都会打印0 2 4 6 8。
func TestSelect2(t *testing.T) {
	// 增加buffer大小会使其输出变得不确定，
	// 因为当buffer既不为满也不为空时，select语句的执行情况就像是抛硬币的行为一样是随机的
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}

// channel的零值是nil。
// 也许会让你觉得比较奇怪，nil的channel有时候也是有一些用处的。
// 因为对一个nil的channel发送和接收操作会永远阻塞，在select语句中操作nil的channel永远都不会被select到。
