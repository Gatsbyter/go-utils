package function

import (
	"testing"
)

// 参数传递，全部是传值
func TestParamPass(t *testing.T) {
	a := 100
	p := &a
	t.Logf("pointer: %p, value: %d\n", p, a)
	changeX(p)
	printPtr(p, t)
	changeY(&p)
	printPtr(p, t)
}

func printPtr(x *int, t *testing.T) {
	t.Logf("pointer: %p, value: %d\n", x, *x)
}

func changeX(x *int) {
	*x++
	//,*x = (*x) + 1
}

func changeY(y **int) {
	**y = (**y) + 1
}
