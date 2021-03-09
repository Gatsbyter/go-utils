package L_io

import (
	"fmt"
	"testing"
)

func TestFmt(t *testing.T) {
	p := &Person{"kema", 21, 0}

	fmt.Println(p)
	fmt.Printf("%L\n", p)
	fmt.Printf("%#v\n", p)
}