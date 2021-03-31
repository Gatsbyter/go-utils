package L_string

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

// len一个字符串 得到的是char[]的长度
// range一个字符串 次数是rune[]的长度
func TestString(t *testing.T) {
	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i, ch := range s {
		fmt.Printf("%c  ---  %c\n", s[i], ch)
	}

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	/*---------------------------------这两个是一个效果-------------------------------------*/

	n1 := 0
	for range s {
		n1++
	}

	n2 := utf8.RuneCountInString(s)

	fmt.Println(n1, n2)
}

func TestRune(t *testing.T) {
	//s := "プログラム"
	s := "hello, 世界"
	// 在第一个Printf中的% x参数用于在每个十六进制数字前插入一个空格
	fmt.Printf("% x\n", s) // 68 65 6c 6c 6f 2c 20 e4 b8 96 e7 95 8c
	// string 转 rune 切片
	r := []rune(s)
	fmt.Printf("%x\n", r) // [68 65 6c 6c 6f 2c 20 4e16 754c]
	// rune切片 转 string
	fmt.Println(string(r))
	// 将一个整数转型为字符串意思是生成以只包含对应Unicode码点字符的UTF8字符串
	fmt.Println(string(rune(65)))     // "A"
	fmt.Println(string(rune(0x4eac))) // "京"
	// 如果对应码点的字符是无效的，则用\uFFFD无效字符作为替换
	fmt.Println(string(rune(1234567))) // "�"

}

//标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。
//
//strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。
//
//bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。
//⚠️因为字符串是只读的，因此逐步构建字符串会导致很多分配和复制。
//在这种情况下，使用bytes.Buffer类型将会更有效，稍后我们将展示。
//
//strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。
//
//unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。
//每个函数有一个单一的rune类型的参数，然后返回一个布尔值。
//而像ToUpper和ToLower之类的转换函数将用于rune字符的大小写转换。
//所有的这些函数都是遵循Unicode标准定义的字母、数字等分类规范。
//strings包也有类似的函数，它们是ToUpper和ToLower，将原始字符串的每个字符都做相应的转换，然后返回新的字符串。
