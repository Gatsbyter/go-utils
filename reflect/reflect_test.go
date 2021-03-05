package reflect

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

// 能够在运行时更新变量和检查它们的值、调用它们的方法和它们支持的内在操作，而不需要在编译时就知道这些变量的具体类型。
// 这种机制被称为反射

// 反射可以探知当前包和外包的非导出结构成员
// 不能对非导出字段直接进行设置操作  不论当前包还是外包

// typeOf 得到类型
func TestReflect1(t *testing.T) {
	// 因为 reflect.TypeOf 返回的是一个动态类型的接口值, 它总是返回具体的类型.
	// 因此, 下面的代码将打印 "*os.File" 而不是 "io.Writer"
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // *os.File

	// %T 参数, 内部使用 reflect.TypeOf 来输出
	fmt.Printf("%T\n", w) // *os.File
}

func TestReflect2(t *testing.T) {
	// reflect.ValueOf 的逆操作是 reflect.Value.Interface 方法.
	//m它返回一个 interface{} 类型，装载着与 reflect.Value 相同的具体值
	v := reflect.ValueOf(3) // reflect.Value类型
	x := v.Interface()      // interface{}类型
	i := x.(int)
	fmt.Printf("%d\n", i) // "3"
}

func TestAny(t *testing.T) {
	x1 := 123
	x2 := 12.34
	x3 := "xxxx"
	x4 := true

	fmt.Println(Any(x1))
	fmt.Println(Any(x2))
	fmt.Println(Any(x3))
	fmt.Println(Any(x4))
}
