package L_map

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	// go 里面的 . 操作符 即是 . 也是 ->

	// ⚠️ 如果结构体成员名字是以大写字母开头的，那么该成员就是导出的；
	// 这是Go语言导出规则决定的。一个结构体可能同时包含导出和未导出的成员
	type T1 struct {
		a, b int // 这俩变量在其他包里 就没法访问到
	} // go 就是这样 任何小写的变量 只能在自己的package里访问到

	// 这几种方式定义都可以
	/*  1  */
	t1 := &T1{1, 2}
	/*  2  */
	t1 = &T1{
		a: 1,
		b: 2,
	}
	/*  3  */
	t1 = new(T1) // 这里只分配了内存 都是0
	*t1 = T1{1, 2}

	fmt.Printf("%#v", *t1)

	// ⚠️ 如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的
	// 那样的话两个结构体将可以使用==或!=运算符进行比较
}

func TestEmbed(t *testing.T) {
	type Embed struct {
		A int
		B int
	}

	type Out struct {
		Embed
		C int
		D int
	}

	o := Out{
		Embed: Embed{
			A: 1,
			B: 2,
		},
		C: 3,
		D: 4,
	}

	// 定义和打印会有点烦
	fmt.Printf("%#v\n", o)
	// 可以直接访问
	fmt.Println(o.A, o.B, o.C, o.D)
}
