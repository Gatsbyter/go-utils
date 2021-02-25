package L_function

import "testing"

// 测试函数返回局部变量指针，安全的！！！
func TestFuncRetPointer(t *testing.T) {
	a := test()
	t.Log(a, *a)
}

func test() *int {
	a := 100
	return &a
}
