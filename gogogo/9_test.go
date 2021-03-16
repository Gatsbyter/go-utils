package gogogo

import (
	"fmt"
	"testing"
)

// increaseA() 的返回参数是匿名，increaseB() 是具名
func Test9(t *testing.T) {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

// defer可以读取有名返回值
func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}
