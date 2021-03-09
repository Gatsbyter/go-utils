package L_io

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// io 包为 I/O 原语提供了基本的接口。它主要包装了这些原语的已有实现。
//
// 由于这些被接口包装的I/O原语是由不同的低级操作实现，因此，在另有声明之前不该假定它们的并发执行是安全的。
//
// 在 io 包中最重要的是两个接口：Reader 和 Writer 接口。
// 各种 IO 包，都跟这两个接口有关，也就是说，只要满足这两个接口，它就可以使用 IO 包的功能。
// ⚠️参考隔壁 io.go

func TestReadFrom(t *testing.T) {
	// 从标准输入读取
	data, err := ReadFrom(os.Stdin, 11)
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(data)

	// 从普通文件读取，其中 file 是 os.File 的实例
	file, err := os.Open("./io_test.go")
	if err != nil {
		panic(err)
	}

	data, err = ReadFrom(file, 9)
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Println(string(data))
}

func TestWrite(t *testing.T) {
	fmt.Println()  // 跳转进去看看
}

func TestReadAt(t *testing.T) {
	reader := strings.NewReader("hello 你好呀")
	p := make([]byte, 7)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)
}

func TestWriteAt(t *testing.T) {
	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n, err := file.WriteString("hello 这里是多余")
	if err != nil {
		panic(err)
	}

	n, err = file.WriteAt([]byte("world world world world"), 6)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
}

func TestReadFrom2(t *testing.T) {
	// ioutil.ReadFile() // 跳进去看看这个

	file, err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(os.Stdout)
	n, err := writer.ReadFrom(file)
	if err != nil {
		panic(err)
	}

	writer.Flush()

	fmt.Println()
	fmt.Println(n)
}

func TestWriteTo(t *testing.T) {
	reader := bytes.NewReader([]byte("Go语言中文网"))

	n, err := reader.WriteTo(os.Stdout)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println(n)
}

func TestSeek(t *testing.T) {
	reader := strings.NewReader("Go语言中文网")
	n, err := reader.Seek(-6, io.SeekEnd)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)

	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
}