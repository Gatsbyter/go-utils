package L_function

import "testing"

// ⚠️在声明方法时，如果一个类型名本身是一个指针的话，是不允许其出现在接收器中的
// type A *int
// func (a A) do() {}  // 这里编译不通过

type Point struct {
	a, b int
}

// 下面的这个p 叫receiver 接收器
func (p Point) Do1()  {} // 传值 不会改变p 但会有一次拷贝
func (p *Point) Do2() {} // 传指针 会改变p

func TestFunction(t *testing.T) {
	// 这个没有问题
	Point{1, 2}.Do1()

	// 这里会编译不过
	// 因为临时变量的内存地址就无法获取得到
	// 我们不能通过一个无法取到地址的接收器来调用指针方法，
	// Point{1, 2}.Do2()

	p1 := Point{1, 2}
	p2 := &Point{1, 2}

	// 都没问题 会隐式转换
	p1.Do1()
	p1.Do2()

	// 都没问题
	p2.Do1()
	p2.Do2() // 方法表达式

	// ⚠️ 这他妈都行
	D1 := p1.Do1 // 方法值  // gin里面传api就是用的这玩意
	D1()
}
