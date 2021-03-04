package routine

import "testing"

// ⚠️
// 两个相同类型的channel可以使用==运算符比较
// 如果两个channel引用的是相同的对象，那么比较的结果为真
// 一个channel也可以和nil进行比较。
func TestEqual(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := ch1

	t.Log(ch1 == ch2) // false
	t.Log(ch1 == ch3) // true
}
