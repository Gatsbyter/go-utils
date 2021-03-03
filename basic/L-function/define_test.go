package L_function

import (
	"fmt"
	"testing"
)

func TestDefine(t *testing.T) {
	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"
}

// 这四种定义方式 都是OK的 类型也是一样的
func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

// ⚠️如果一个函数将所有的返回值都显示的变量名，那么该函数的return语句可以省略操作数。这称之为bare return
// 不宜过度使用bare return
