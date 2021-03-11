package L_bit

import (
	"fmt"
	"testing"
)

func TestBit(t *testing.T) {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010"
	fmt.Printf("%08b\n", y) // "00000110"

	fmt.Printf("%08b\n", x&y)  // "00000010"
	fmt.Printf("%08b\n", x|y)  // "00100110"
	fmt.Printf("%08b\n", x^y)  // "00100100", 不同的位是1
	fmt.Printf("%08b\n", x&^y) // "00100000", x是1 y是0的位是1

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}
