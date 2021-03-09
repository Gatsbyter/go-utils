package L_io

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// 虽然 io 包提供了不少类型、方法和函数，但有时候使用起来不是那么方便。
// 比如读取一个文件中的所有内容。
// 为此，标准库中提供了一些常用、方便的IO操作函数。

// ioutil.ReadAll
func TestReadAll(t *testing.T) {
	file, err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}

	// 1.16 之后 这个和 io.ReadAll等价
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func TestReadDir(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(pwd, "../../")
	fmt.Printf("target path : %s\n\n", path)

	listAll(path, 0)
}

// ReadFile 从 filename 指定的文件中读取数据并返回文件的内容。
// 成功的调用返回的err 为 nil 而非 EOF。
// 因为本函数定义为读取整个文件，它不会将读取返回的 EOF 视为应报告的错误。(同 ReadAll )
// 
// ReadFile 会先判断文件的大小，给 bytes.Buffer 一个预定义容量，避免额外分配内存。
func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("writeAt.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

// WriteFile 将data写入filename文件中，
// 当文件不存在时会根据perm指定的权限进行创建一个,文件存在时会先清空文件内容。
// 对于 perm 参数，我们一般可以指定为：0666，具体含义 os 包中讲解。
func TestWriteFile(t *testing.T) {
	err := ioutil.WriteFile("writeFile.txt", []byte("I'm writing this file"), 0666)
	if err != nil {
		panic(err)
	}
}

// 操作系统中一般都会提供临时目录，比如 linux 下的 /tmp 目录（通过 os.TempDir() 可以获取到)。
// 有时候，我们自己需要创建临时目录，比如 Go 工具链源码中（src/cmd/go/build.go）
//
// 通过 TempDir 创建一个临时目录，用于存放编译过程的临时文件
func TestTempDir(t *testing.T) {
	// 第一个参数如果为空，表明在系统默认的临时目录（ os.TempDir ）中创建临时目录；
	// 第二个参数指定临时目录名的前缀，该函数返回临时目录的路径。
	work, err := ioutil.TempDir("", "go-utils-test")
	if err != nil {
		panic(err)
	}

	fmt.Println(work)

	err = os.Remove(work)
	if err != nil {
		panic(err)
	}
}

func TestTempFile(t *testing.T)  {
	file, err := ioutil.TempFile("", "go-utils-file")
	if err != nil {
		panic(err)
	}

	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	fmt.Println(file.Name())
}