package L_io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

// Scanner 可以更容易的处理如按行读取输入序列或空格分隔单词等
// 它终结了如输入一个很长的有问题的行这样的输入错误，并且提供了简单的默认行为：基于行的输入，每行都剔除分隔标识
func TestScanner(t *testing.T) {
	// ⚠️ 这个要用main测试 别用test
	scanner := bufio.NewScanner(os.Stdin) // 像这些里面都定义了 SplitFunc 定义后就可以调Scan 可以手动覆盖(用Split

	for scanner.Scan() { // Scan 方法 该方法好比 iterator 中的 Next 方法，它用于将 Scanner 获取下一个 token
		fmt.Println(scanner.Text()) // Bytes 和 Text 方法 这两个方法的行为一致，都是返回最近的 token，该方法应该在 Scan 调用后调用
	}

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// SplitFunc 输入的行为可以通过一个函数控制，来控制输入的每个部分，但是对于复杂的问题或持续传递错误的，可能还是需要原有接口。
	// 参考粘包
}

// 一行行读取文件
func TestScanLine(t *testing.T) {
	file, err := os.Create("scanner.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, _ = file.WriteString("first line\nsecond line\nthird line")

	// 将文件 offset 设置到文件开头
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	// 这儿默认按行split
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
