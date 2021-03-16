package gogogo

import (
	"fmt"
	"testing"
)

// 截断slice cap默认和原先一样
func Test6(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[3:4:4] // 第三个是cap
	fmt.Println(b[0], len(b), cap(b))
}
