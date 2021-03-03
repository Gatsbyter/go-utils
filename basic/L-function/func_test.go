package L_function

import (
	"fmt"
	"strings"
	"testing"
)

// 测试函数返回局部变量指针，安全的！！！
func TestFuncRetPointer(t *testing.T) {
	a := test()
	t.Log(a, *a)
}

func test() *int {
	a := 100
	return &a
}

func TestNilFunc(t *testing.T) {
	var f func(int) int
	f(3) // 此处f的值为nil, 会引起panic错误

	// 所以这种用之前先判断一下
	if f != nil {
		f(3)
	}
}

// ⚠️函数值之间是不可比较的，也不能用函数值作为map的key

func TestMapFunc(t *testing.T) {
	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
}

func add1(r rune) rune { return r + 1 }
