package gogogo

import (
	"fmt"
	"testing"
)

// ⚠️ nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量

const (
	x = iota
	_ // 占位也会++
	y
	z = "zz" // 出现其他类型 后面跟上
	k
	p = iota // iota重现江湖 回来
	q = iota
)

const a = iota // 出现const 归0

func Test4(t *testing.T) {
	fmt.Println(x, y, z, k, p, q, a)
}
