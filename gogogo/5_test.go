package gogogo

import (
	"fmt"
	"testing"
)

// map读写问题
func Test5(t *testing.T) {
	var m map[int]int
	_ = m[1] // 读取不会报错
	m[1] = 1 // 赋值就会报错了

	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}

func hello(num ...int) {
	num[0] = 18
}
