package func_learning

import "testing"

// 来深入浅出一下闭包
// 通过指针饮用上下文环境变量，导致其生命周期延长
func TestClosure(t *testing.T) {
	a := 100
	t.Log(&a, a)
	f := closure(a, t)
	f()
}

func closure(x int, t *testing.T) func() {
	t.Log(&x, x)

	return func() {
		t.Log(&x, x)
	}
}

// 闭包只存指针
// 所以闭包的环境变量很关键，很容易搞出错
func TestClosure2(t *testing.T) {
	for _, f := range closure2(t) {
		f()
	}
}

func closure2(t *testing.T) []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		t.Logf("nimabi-- %p -- %d", &i, i)
		s = append(s, func() {
			t.Logf("fuck-- %p -- %d", &i, i)
		})
	}

	return s
}
