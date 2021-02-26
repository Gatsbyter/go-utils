package L_slice

import (
	"fmt"
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
