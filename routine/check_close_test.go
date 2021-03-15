package routine

import (
	"fmt"
	"testing"
	"time"
)

// ⚠️ go语言无法即时准确地判断channel是否关闭

func TestCheckClose(t *testing.T) {
	ch := make(chan int, 1)

	go func() {
		ch <- 1 //
	}()

	time.Sleep(time.Second)
	close(ch) //

	fmt.Println(IsClose1(ch))
}

// 扯淡
func IsClose1(ch <-chan int) bool {
	_, ok := <-ch
	return !ok
}

// 扯淡
func IsClose2(ch <-chan int) bool {
	select {
	case <-ch: // 关了 或者有数据 会进这个
		return true
	default: // 没关 并且没数据 会进这个
	}

	return false
}
