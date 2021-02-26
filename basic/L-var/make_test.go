package L_var

import (
	"fmt"
	"reflect"
	"testing"
)

// make 只能初始化slice、map、chan
// new和make都是在堆上分配内存空间，new只开辟空间并返回指针，make开辟空间并初始化返回类型

func TestNew(t *testing.T) {
	aa := new([]int)
	bb := make([]int, 0)

	// 返回类型不一样
	fmt.Printf("%v\n", reflect.TypeOf(aa))
	fmt.Printf("%v\n", reflect.TypeOf(bb))

	// new出来的全是0啥也没 只有内存空间
	// make还会给你弄一弄
	fmt.Printf("%v\n", *aa == nil)
	fmt.Printf("%v\n", bb == nil)
}
