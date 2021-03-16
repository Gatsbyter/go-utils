package gogogo

import (
	"fmt"
	"testing"
)

// 不同类型相加问题
func Test7(t *testing.T) {
	a := 5
	b := 1.5
	_, _ = a, b
	// fmt.Println(a + b) // 报错

	fmt.Println(5 + 1.5) // 这就不会
}
