package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

type A struct {
	x bool
	y float64
	z int16
} // 3 words
type B struct {
	float64
	int16
	bool
} // 2 words
type C struct {
	bool
	int16
	float64
} // 2 words

func TestStruct(t *testing.T) {
	fmt.Println("A      : ", unsafe.Sizeof(A{}))
	fmt.Println("B      : ", unsafe.Sizeof(B{}))
	fmt.Println("C      : ", unsafe.Sizeof(C{}))
}

func TestAlign(t *testing.T) {
	a := A{}
	fmt.Printf("size: %d, align: %d\n", unsafe.Sizeof(a), unsafe.Alignof(a))

	// unsafe.Alignof 函数返回对应参数的类型需要对齐的倍数
	// unsafe.Offsetof 函数返回 结构体字段相对于 结构体起始地址的偏移量
	fmt.Printf("size: %d, align: %d, offset: %d\n", unsafe.Sizeof(a.x), unsafe.Alignof(a.x), unsafe.Offsetof(a.x))
	fmt.Printf("size: %d, align: %d, offset: %d\n", unsafe.Sizeof(a.y), unsafe.Alignof(a.y), unsafe.Offsetof(a.y))
	fmt.Printf("size: %d, align: %d, offset: %d\n", unsafe.Sizeof(a.z), unsafe.Alignof(a.z), unsafe.Offsetof(a.z))
}
