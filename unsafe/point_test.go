package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

// 大多数指针类型会写成*T，表示是“一个指向T类型变量的指针”。
// unsafe.Pointer是特别定义的一种指针类型（类似C语言中的void*类型的指针）
// 它可以包含任意类型变量的地址。
//
// 当然，我们不可以直接通过*p来获取unsafe.Pointer指针指向的真实变量的值，因为我们并不知道变量的具体类型。
// 和普通指针一样，unsafe.Pointer指针也是可以比较的，并且支持和nil常量比较判断是否为空指针。
//
// 一个普通的*T类型指针可以被转化为unsafe.Pointer类型指针，
// 并且一个unsafe.Pointer类型指针也可以被转回普通的指针，被转回普通的指针类型并不需要和原始的*T类型相同。
// 通过将*float64类型指针转化为*uint64类型指针，我们可以查看一个浮点数变量的位模式。

func TestPointer(t *testing.T) {
	var f float64
	var u uint64
	f = 1.0

	u = *(*uint64)(unsafe.Pointer(&f))
	fmt.Printf("0x%x", u)
}

// 这个牛逼啊 感觉来到了c++
func TestPointer2(t *testing.T) {
	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&X)) + unsafe.Offsetof(X.b)))
	*pb = 42
	fmt.Println(X.b) // "42"

	// unsafe.Pointer 和 uintptr可以相互转换
	// unsafe.Pointer可以和 *T转换  uintptr不行
	// uintptr可以加减 unsafe.Pointer不行
}

// ⚠️和test2类似 但是这样的话 可能会出错
func TestPointer3(t *testing.T) {
	tmp := uintptr(unsafe.Pointer(&X)) + unsafe.Offsetof(X.b)
	pb := (*int16)(unsafe.Pointer(tmp))
	*pb = 42
	fmt.Println(X.b)

	// 产生错误的原因:
	// 有时候垃圾回收器会移动一些变量以降低内存碎片等问题。这类垃圾回收器被称为移动GC。
	// 当一个变量被移动，所有的保存该变量旧地址的指针必须同时被更新为变量移动后的新地址。
	// 从垃圾收集器的视角来看，一个unsafe.Pointer是一个指向变量的指针，因此当变量被移动时对应的指针也必须被更新；
	// 但是uintptr类型的临时变量只是一个普通的数字，所以其值不应该被改变。
	//
	// 上面错误的代码因为引入一个非指针的临时变量tmp，导致垃圾收集器无法正确识别这个是一个指向变量x的指针。
	// 当第二个语句执行时，变量x可能已经被转移，这时候临时变量tmp也就不再是现在的&x.b地址。
	// 第三个向之前无效地址空间的赋值语句将彻底摧毁整个程序！
}

// 错误!
func TestPointer4(t *testing.T) {
	pT := uintptr(unsafe.Pointer(new(int)))
	fmt.Printf("%x\n", pT)

	// 这里并没有指针引用new新创建的变量，因此该语句执行完成之后，垃圾收集器有权马上回收其内存空间，所以返回的pT将是无效的地址。
	// 但是目前好像没这个问题 whatever
	//
	// 虽然目前的Go语言实现还没有使用移动GC，但这不该是编写错误代码侥幸的理由：
	// 当前的Go语言实现已经有移动变量的场景，goroutine的栈是根据需要动态增长的。
	// 当发生栈动态增长的时候，原来栈中的所有变量可能需要被移动到新的更大的栈中，
	// 所以我们并不能确保变量的地址在整个使用周期内是不变的。

	// 强烈建议按照最坏的方式处理。
	// 将所有包含变量地址的uintptr类型变量当作BUG处理，同时减少不必要的unsafe.Pointer类型到uintptr类型的转换。
	// 在第一个例子中，有三个转换——字段偏移量到uintptr的转换和转回unsafe.Pointer类型的操作——所有的转换全在一个表达式完成

	// 当调用一个库函数，并且返回的是uintptr类型地址时
	// 比如反射包中的相关函数，返回的结果应该立即转换为unsafe.Pointer以确保指针指向的是相同的变量。
	//
	// 普通方法实现的函数尽量不要返回uintptr类型。
	// reflect包和unsafe包一样都是采用特殊技术实现的，编译器可能给它们开了后门
}

var X struct {
	a bool
	b int16
	c []int
}
