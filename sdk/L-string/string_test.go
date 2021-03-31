package L_string

import (
	"fmt"
	"strings"
	"testing"
)

// 字符串切割
func TestString(t *testing.T) {
	s1 := strings.Fields("hello world")
	p(s1)

	s2 := strings.FieldsFunc("哈哈哈淦哈哈哈淦哈哈哈", func(r rune) bool {
		if r == '淦' {
			return true
		}
		return false
	})
	p(s2)
}

func p(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}
