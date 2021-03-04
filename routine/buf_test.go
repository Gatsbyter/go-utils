package routine

import (
	"fmt"
	"testing"
)

// 向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。
// 如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作而释放了新的队列空间。
// 相反，如果channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素
func TestBuf(t *testing.T) {
	ch := make(chan int, 3)

	fmt.Println(cap(ch)) // 3
	fmt.Println(len(ch)) // 0

	ch <- 100
	ch <- 200

	fmt.Println(cap(ch)) // 3
	fmt.Println(len(ch)) // 2

	fmt.Println(<-ch)
	fmt.Println(cap(ch)) // 3
	fmt.Println(len(ch)) // 1

	ch <- 300
	ch <- 400
	// ch <- 500 // 这个再加就会报错了 buf满了 又没人收 死锁
	// 如果再来一个routine就不会死锁了

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<- ch)  // 这个再加也会报错了 buf空了 死锁
}
