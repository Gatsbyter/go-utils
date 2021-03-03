package L_function

import (
	"fmt"
	"testing"
)

// 来深入浅出一下闭包
// 通过指针饮用上下文环境变量，导致其生命周期延长
func TestClosure(t *testing.T) {
	a := 100
	fmt.Println(&a, a)
	f := closure(a)
	f()
}

func closure(x int) func() {
	fmt.Println(&x, x)

	return func() {
		fmt.Println(&x, x)
	}
}

// 闭包只存指针
// 所以闭包的环境变量很关键，很容易搞出错
func TestClosure2(t *testing.T) {
	for _, f := range closure2() {
		f()
	}
}

func closure2() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		fmt.Printf("nimabi-- %p -- %d\n", &i, i)
		s = append(s, func() {
			fmt.Printf("fuck-- %p -- %d\n", &i, i)
		})
	}

	return s
}

// ⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️
// squares的例子证明，函数值不仅仅是一串代码，还记录了状态。
// 在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，
// 这意味着匿名函数和squares中，存在变量引用。
// 这就是函数值属于引用类型和函数值不可比较的原因
func TestClosure3(t *testing.T) {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"

	// 通过这个例子，我们看到变量的生命周期不由它的作用域决定：squares返回后，变量x仍然隐式的存在于f中。
}

// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
