package gogogo

import (
	"fmt"
	"testing"
)

// struct只有相同的才可以比较 成员顺序不同都是不同
// 且只能比较相等 不能比大小
//
// map相同也不能比较
// slice map func 都是不能比较的 但是可以和nil比较
func Test2(t *testing.T) {
	a := struct {
		age  int
		name string
	}{age: 10, name: "haha"}

	b := struct {
		age  int
		name string
	}{age: 10, name: "haha"}

	fmt.Println(a == b)

	m1 := map[string]int{
		"a": 1,
	}

	m2 := map[string]int{
		"a": 1,
	}

	_, _ = m1, m2
	// fmt.Println(m1 == m2)
}
