package routine

import (
	"fmt"
	"testing"
)

// 两种判断channel是否关闭的办法
// 推荐2 完美运行
func TestCheckClose(t *testing.T) {
	ch := make(chan int)

	close(ch)

	fmt.Println(IsClose2(ch))
}

// channel关闭了 是可以判断的，没有影响
// 没关闭的话，单routine容易死锁
func IsClose1(ch <-chan int) bool {
	_, ok := <-ch
	return !ok
}

func IsClose2(ch <-chan int) bool {
	select {
	case <-ch: // 关了 会进这个
		return true
	default: // 没关 会进这个
	}

	return false
}
