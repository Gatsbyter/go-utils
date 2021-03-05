package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

// 按字节打印一个uint64 也可以测试大小端编码
func TestByte(t *testing.T) {
	var arr [8]byte
	var u uint64 = 0x1122334455667788

	for i, _ := range arr {
		arr[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&u)) + uintptr(i)))
	}

	fmt.Printf("%x\n", u)
	fmt.Printf("%#v\n", arr)
}
