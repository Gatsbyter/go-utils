package L_io

import (
	"bytes"
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

// 根据 Go 语言中实现接口的定义，一个类型只要有 String() string 方法，我们就说它实现了 Stringer 接口。
// 而在本节开始已经说到，如果格式化输出某种类型的值，只要它实现了 String() 方法，那么会调用 String() 方法进行处理。
func (p *Person) String() string {
	buffer := bytes.NewBufferString("This is ")
	buffer.WriteString(p.Name + ", ")
	if p.Sex == 0 {
		buffer.WriteString("He ")
	} else {
		buffer.WriteString("She ")
	}

	buffer.WriteString("is ")
	buffer.WriteString(strconv.Itoa(p.Age))
	buffer.WriteString(" years old.")
	return buffer.String()
}

// Formatter 接口由带有定制的格式化器的值所实现。
// Format 的实现可调用 Sprintf 或 Fprintf(f) 等函数来生成其输出。
func (p *Person) Format(f fmt.State, c rune) {
	if c == 'L' {
		f.Write([]byte(p.String()))
		f.Write([]byte(" Person has three fields."))
	} else {
		// 没有此句，会导致 fmt.Printf("%s", p) 啥也不输出
		f.Write([]byte(fmt.Sprintln(p.String())))
	}
}

// 1. fmt.State 是一个接口。由于 Format 方法是被 fmt 包调用的，它内部会实例化好一个 fmt.State 接口的实例，我们不需要关心该接口；
//
// 2. 可以实现自定义占位符，同时 fmt 包中和类型相对应的预定义占位符会无效。因此例子中 Format 的实现加上了 else 子句
//
// 3. 实现了 Formatter 接口，相应的 Stringer 接口不起作用。
//    但实现了 Formatter 接口的类型应该实现 Stringer 接口，这样方便在 Format 方法中调用 String() 方法
//
// 4. Format 方法的第二个参数是占位符中%后的字母（有精度和宽度会被忽略，只保留字母）；
//
// 一般地，我们不需要实现 Formatter 接口。如果对 Formatter 接口的实现感兴趣，可以看看标准库 math/big 包中 Int 类型的 Formatter 接口实现。


// 该接口定义了类型的Go语法格式。用于打印(Printf)格式化占位符为 %#v 的值。
func (p *Person) GoString() string {
	return "&Person{Name is "+p.Name+", Age is "+strconv.Itoa(p.Age)+", Sex is "+strconv.Itoa(p.Sex)+"}"
}
