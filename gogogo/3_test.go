package gogogo

import (
	"fmt"
	"testing"
)

// append可能会开辟新空间
func Test3(t *testing.T) {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)
}

func ap(a []int) {
	a = append(a, 10)
}

func app(a []int) {
	a[0] = 1
}
