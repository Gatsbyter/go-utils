package L_map

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// 初始化的方法
	var m1 map[string]int // nil
	m2 := map[string]int{}
	m3 := map[string]int(nil) // nil
	m4 := make(map[string]int)

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)

	fmt.Println(m1 == nil)
	fmt.Println(m2 == nil)
	fmt.Println(m3 == nil)
	fmt.Println(m4 == nil)

	// ⚠️ map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似。
	// 但是向一个nil值的map存入元素将导致一个panic异常
	// m1["carol"] = 21

	// ⚠️ map[T]bool 可以实现一个类似set
}

// 操作map
func TestOp(t *testing.T) {
	// 这两种方式是等价的
	ages0 := make(map[string]int)
	ages0["alice"] = 31
	ages0["charlie"] = 34

	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	fmt.Println(len(ages))

	age, ok := ages["alice"]
	fmt.Println(age, ok)

	ages["alice"] = 32

	age, ok = ages["alice"]
	fmt.Println(age, ok)

	delete(ages, "alice")

	age, ok = ages["alice"]
	fmt.Println(age, ok)

	// ⚠️不能对map的元素进行取址操作
	// 禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效
	// _ = &ages["bob"]
}

// ⚠️ Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。
// 在实践中，遍历的顺序是随机的，每一次遍历的顺序都不相同
// 如果要按顺序遍历key/value对，我们必须显式地对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序

// ⚠️ 和slice一样，map之间也不能进行相等比较 唯一的例外是和nil进行比较。
// 要判断两个map是否包含相同的key和value，我们必须通过一个循环实现
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	// 用!ok来区分元素不存在，与元素存在但为0的
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

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
