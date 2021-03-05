package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

// 其中a对应的变量不可取地址。因为a中的值仅仅是整数2的拷贝副本。
// b中的值也同样不可取地址。
// c中的值还是不可取地址，它只是一个指针&x的拷贝。
// 实际上，所有通过reflect.ValueOf(x)返回的reflect.Value都是不可取地址的。
// 但是对于d，它是c的解引用方式生成的，指向另一个变量，因此是可取地址的。
// 我们可以通过调用reflect.ValueOf(&x).Elem()，来获取任意变量x对应的可取地址的Value
func TestValue(t *testing.T) {
	x := 2                   // value   type    variable?
	a := reflect.ValueOf(2)  // 2       int     no
	b := reflect.ValueOf(x)  // 2       int     no
	c := reflect.ValueOf(&x) // &x      *int    no
	d := c.Elem()            // 2       int     yes (x)

	fmt.Println(a.CanAddr(), a.CanSet()) // "false"
	fmt.Println(b.CanAddr(), b.CanSet()) // "false"
	fmt.Println(c.CanAddr(), c.CanSet()) // "false"
	fmt.Println(d.CanAddr(), d.CanSet()) // "true"
	fmt.Println(a, b, c, d)
}

// 要从变量对应的可取地址的reflect.Value来访问变量需要三个步骤。
// 第一步是调用Addr()方法，它返回一个Value，里面保存了指向变量的指针。
// 然后是在Value上调用Interface()方法，也就是返回一个interface{}，里面包含指向变量的指针。
// 最后，如果我们知道变量的类型，我们可以使用类型的断言机制将得到的interface{}类型的接口强制转为普通的类型指针。
// 这样我们就可以通过这个普通指针来更新变量了
func TestAddr(t *testing.T) {
	x := 2
	d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	px := d.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)                    // "3"
}

// 不使用指针，而是通过调用可取地址的reflect.Value的reflect.Value.Set方法来更新对于的值
func TestAddr2(t *testing.T) {
	x := 2
	d := reflect.ValueOf(&x).Elem() // d refers to the variable x
	d.Set(reflect.ValueOf(4))
	fmt.Println(x) // "4"

	// Set方法将在运行时执行和编译时进行类似的可赋值性约束的检查。
	// 以上代码，变量和值都是int类型，但是如果变量是int64类型，那么程序将抛出一个panic异常，
	// 所以关键问题是要确保改类型的变量可以接受对应的值：
	// d.Set(reflect.ValueOf(int64(5)))   // 这个会报错
}

// 对一个不可取地址的reflect.Value调用Set方法也会导致panic异常
func TestAddr3(t *testing.T) {
	x := 2
	b := reflect.ValueOf(x)
	b.Set(reflect.ValueOf(3)) // panic: Set using unaddressable value
}
