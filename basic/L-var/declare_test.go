package L_var

import "testing"

func TestDeclare(t *testing.T) {
	// 变量的申明方式
	var i, j, k int
	var b, f, s = true, 2.3, "ok"
	a1, a2, a3 := 1, 2, "fuck"

	t.Log(i, j, k)
	t.Log(b, f, s)
	t.Log(a1, a2, a3)

	// 这是空指针
	var p *int
	t.Logf("%p", p)
}

func TestPointer(t *testing.T) {
	a := 1
	pointer(&a)
	t.Log(a)
}

func pointer(p *int) {
	if p == nil {
		return
	}

	*p += 1
}
