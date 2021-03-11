package L_slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	var arr1 [6]int
	arr2 := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(arr1)
	// arr 的 len和cap相同
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))
	fmt.Println(arr2)

	// arr 无法append
}

func TestDiff(t *testing.T) {
	// slice本质是array元素的引用，官方把slice叫做underlying array
	// slice传递的是地址，效果和指针相同，而array是复制元素
	arr := [6]int{11, 22, 33, 44, 55, 66}
	sli := []int{1, 2, 3, 4, 5, 6}

	func(x [6]int) {
		x[2] = 333
	}(arr)

	func(x []int) {
		x[2] = 333
	}(sli)

	fmt.Println("arr:", arr)
	fmt.Println("sli:", sli)
}

func TestInitArr(t *testing.T) {
	type Currency int

	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)

	// 可以这样初始化
	// ⚠️ 三点初始化 是数组 不是slice
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println(RMB, symbol[RMB])

	// 也可以这样初始化 其他位置都是0
	r := [...]int{99: -1}
	fmt.Println(r)
}

func TestCompare(t *testing.T) {
	// 只有当两个数组的所有元素都是相等的时候数组才是相等
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	// d := [3]int{1, 2}
	// 直接编译错误 [2]int == [3]int

	// 和数组不同的是，slice之间不能比较
	// 因此我们不能使用==操作符来判断两个slice是否含有全部相等元素。
	// 不过标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相等（[]byte）
	// 但是对于其他类型的slice，我们必须自己展开每个元素进行比较
}

func TestType(t *testing.T) {
	// 数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型
	a1 := [3]int{1, 2, 3}   // 数组
	a2 := [...]int{1, 2, 3} // 数组
	a3 := []int{1, 2, 3}    // 切片

	fmt.Println(reflect.TypeOf(a1))
	fmt.Println(reflect.TypeOf(a1[:])) // 切片
	fmt.Println(reflect.TypeOf(a2))
	fmt.Println(reflect.TypeOf(a3))
}

func TestCall(t *testing.T) {
	zero(&[32]byte{})
}

// 这个函数 有很多知识点
func zero(ptr *[32]byte) {
	// 传参 如果直接传数组，那就要把整个数组copy一遍
	// 所以传指针
	for i, _ := range ptr { // 这里 go操作指针不需要像c++ 加*号
		ptr[i] = 0
	}

	fmt.Printf("%v, %[1]T\n", ptr)
	fmt.Printf("%v, %[1]T\n", *ptr)
}
