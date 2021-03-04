package routine

import "testing"

// Channel还支持close操作，用于关闭channel
// 随后对基于该channel的任何 [发送] 操作都将导致panic异常
func TestClose1(t *testing.T) {
	ch := make(chan int, 10)

	ch <- 100

	close(ch)

	ch <- 200 // panic: send on closed channel
}

// 对一个已经被close过的channel进行接收操作依然可以接受到之前已经成功发送的数据
// 如果channel中已经没有数据的话将产生一个零值的数据
func TestClose2(t *testing.T) {
	ch := make(chan int, 10)

	ch <- 100

	t.Log(<-ch)
	// t.Log(<-ch)   // 这儿是不能收的 卡死了

	ch <- 100
	close(ch)

	t.Log(<-ch) // 100
	t.Log(<-ch) // 0  close之后 可以瞎几把收 卡不死
	t.Log(<-ch)
	t.Log(<-ch)
}
