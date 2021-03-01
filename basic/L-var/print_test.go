package L_var

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	type user struct {
		name string
		age  float32
	}

	u := user{
		name: "me",
		age:  13.5,
	}

	/*----------------------------------复杂类型--------------------------------------*/

	fmt.Printf("%v\n", u)  // {me 13.5}
	fmt.Printf("%+v\n", u) // {name:me age:13.5}
	fmt.Printf("%#v\n", u) // L_var.user{name:"me", age:13.5}

	fmt.Printf("%p\n", &u) //指针表示
	fmt.Printf("%T\n", u)  // 输出变量的类型

	/*-----------------------------------整数类型-------------------------------------*/

	fmt.Printf("%b\n", 1024) //二进制表示
	fmt.Printf("%d\n", 10)   //十进制表示
	fmt.Printf("%o\n", 8)    //八进制表示

	fmt.Printf("%x\n", 1223) // 十六进制表示，小写
	fmt.Printf("%X\n", 1223) // 十六进制表示，大写
	fmt.Printf("%q\n", 22)   // 转化为十六进制并附上单引号 '\x16'

	/*----------------------------------字符串类型------------------------------------*/

	fmt.Printf("%s\n", "wqdew")  // 直接输出字符串或者[]byte
	fmt.Printf("%q\n", "dedede") // 双引号括起来的字符串

	fmt.Printf("%c\n", 12345) // 数值对应的 Unicode 编码字符 〹
	fmt.Printf("%U\n", 1233)  // Unicode表示  U+04D1

	/*----------------------------------浮点--------------------------------------*/

	fmt.Printf("%f\n", 12.3456) // 浮点打印

	fmt.Printf("%e\n", 12.345)   // 科学计数法，e表示 1.234500e+01
	fmt.Printf("%E\n", 12.34455) // 科学计数法，E表示 1.234455E+01

	fmt.Printf("%g\n", 12.3456) // 根据实际情况采用%e或%f输出
	fmt.Printf("%G\n", 12.3456) // 根据实际情况采用%E或%f输出

	/*------------------------------------------------------------------------*/

	fmt.Printf("%t\n", true) // 布尔值输出 其实没卵用 %v就行了

	/*------------------------------------------------------------------------*/

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
}
