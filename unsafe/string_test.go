package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

// string的内部结构是这样的
// type stringStruct struct {
//	 str unsafe.Pointer
//	 len int
// }

// string + 的时候 *char 会重新分配内存
// ⚠️ https://my.oschina.net/renhc/blog/3019849

func TestString(t *testing.T) {
	str := "hello"
	fmt.Printf("%p\n", &str)

	sp := uintptr(unsafe.Pointer(&str))
	fmt.Printf("0x%x\n", sp)

	strStart := *(*uintptr)(unsafe.Pointer(sp))
	fmt.Printf("0x%x\n", strStart)

	ch0 := *(*uint8)(unsafe.Pointer(strStart))
	ch1 := *(*uint8)(unsafe.Pointer(strStart + 1))
	ch2 := *(*uint8)(unsafe.Pointer(strStart + 2))
	ch3 := *(*uint8)(unsafe.Pointer(strStart + 3))
	ch4 := *(*uint8)(unsafe.Pointer(strStart + 4))
	fmt.Printf("%c\n", ch0)
	fmt.Printf("%c\n", ch1)
	fmt.Printf("%c\n", ch2)
	fmt.Printf("%c\n", ch3)
	fmt.Printf("%c\n", ch4)

	length := *(*uint64)(unsafe.Pointer(sp + 8))
	fmt.Printf("%d\n", length)

	str += "world"

	sp = uintptr(unsafe.Pointer(&str))
	fmt.Printf("0x%x\n", sp)

	strStart = *(*uintptr)(unsafe.Pointer(sp))
	fmt.Printf("0x%x\n", strStart)

	ch0 = *(*uint8)(unsafe.Pointer(strStart))
	ch1 = *(*uint8)(unsafe.Pointer(strStart + 1))
	ch2 = *(*uint8)(unsafe.Pointer(strStart + 2))
	ch3 = *(*uint8)(unsafe.Pointer(strStart + 3))
	ch4 = *(*uint8)(unsafe.Pointer(strStart + 4))
	fmt.Printf("%c\n", ch0)
	fmt.Printf("%c\n", ch1)
	fmt.Printf("%c\n", ch2)
	fmt.Printf("%c\n", ch3)
	fmt.Printf("%c\n", ch4)

	length = *(*uint64)(unsafe.Pointer(sp + 8))
	fmt.Printf("%d\n", length)
}
