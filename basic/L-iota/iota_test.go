package L_iota

import (
	"fmt"
	"testing"
)

func TestIota(t *testing.T) {
	// 如果是批量声明的常量，除了第一个外其它的常量右边的初始化表达式都可以省略
	// 如果省略初始化表达式则表示使用前面常量的初始化表达式写法，对应的常量类型也一样的
	const (
		a = 1
		b
		c = 2
		d
	)

	fmt.Println(a, b, c, d) // "1 1 2 2"

	// 在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一
	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	const (
		_   = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
		TiB // 1099511627776             (exceeds 1 << 32)
		PiB // 1125899906842624
		EiB // 1152921504606846976
		ZiB // 1180591620717411303424    (exceeds 1 << 64)
		YiB // 1208925819614629174706176
	)

	fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB)

	// Go语言的常量有个不同寻常之处。
	// 虽然一个常量可以有任意有一个确定的基础类型，例如int或float64，或者是类似time.Duration这样命名的基础类型
	// 但是许多常量并没有一个明确的基础类型。
	// 编译器为这些没有明确的基础类型的数字常量提供比基础类型更高精度的算术运算
	// 你可以认为至少有256bit的运算精度。
	//
	// 这里有六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。
	//
	// 通过延迟明确常量的具体类型，无类型的常量不仅可以提供更高的运算精度
	// 而且可以直接用于更多的表达式而不需要显式的类型转换
	// 例如，例子中的ZiB和YiB的值已经超出任何Go语言中整数类型能表达的范围
	// 但是它们依然是合法的常量，而且可以像下面常量表达式依然有效
	// YiB/ZiB是在编译期计算出来的，并且结果常量是1024，是Go语言int变量能有效表示的

	fmt.Println(YiB / ZiB) // "1024"
}
