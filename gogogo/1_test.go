package gogogo

import (
	"testing"
)

// for循环中 v的地址问题
func Test1(t *testing.T) {
	s := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for k, v := range s {
		m[k] = &v
	}

	for k, v := range m {
		println(k, " --> ", *v)
	}
}
