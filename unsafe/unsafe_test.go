package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	// unsafe.Sizeof函数返回操作数在内存中的字节大小，参数可以是任意类型的表达式，但是它并不会对表达式进行求值。
	// 一个Sizeof函数调用是一个对应uintptr类型的常量表达式，因此返回的结果可以用作数组类型的长度大小，或者用作计算其他的常量

	// Sizeof函数返回的大小只包括数据结构中固定的部分，
	// 例如字符串对应结构体中的指针和字符串长度部分，但是并不包含指针指向的字符串的内容

	// 由于地址对齐这个因素，一个聚合类型（结构体或数组）的大小至少是所有字段或元素大小的总和，或者更大因为可能存在内存空洞。
	// 内存空洞是编译器自动添加的没有被使用的内存空间，用于保证后面每个字段或元素的地址相对于结构或数组的开始地址能够合理地对齐
	// 内存空洞可能会存在一些随机数据，可能会对用unsafe包直接操作内存的处理产生影响
	var (
		b      bool
		i      int
		u      uint
		up     uintptr
		ptr    *int
		s      string
		arr    []int
		m      map[string]string
		f      func()
		ch     chan int
		interf interface{}
	)

	// unsafe.Sizeof 返回的事uintptr类型
	fmt.Println("bool      : ", unsafe.Sizeof(b))
	fmt.Println("int       : ", unsafe.Sizeof(i))
	fmt.Println("uint      : ", unsafe.Sizeof(u))
	fmt.Println("uintptr   : ", unsafe.Sizeof(up))
	fmt.Println("*T        : ", unsafe.Sizeof(ptr))
	fmt.Println("string    : ", unsafe.Sizeof(s))
	fmt.Println("[]T       : ", unsafe.Sizeof(arr))
	fmt.Println("map       : ", unsafe.Sizeof(m))
	fmt.Println("func      : ", unsafe.Sizeof(f))
	fmt.Println("chan      : ", unsafe.Sizeof(ch))
	fmt.Println("interface : ", unsafe.Sizeof(interf))
}
