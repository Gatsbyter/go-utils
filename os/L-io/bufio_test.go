package L_io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

// bufio 包实现了缓存IO。
// 它包装了 io.Reader 和 io.Writer 对象，创建了另外的Reader和Writer对象，
// 它们也实现了 io.Reader 和 io.Writer 接口，不过它们是有缓存的

// 当然是为了提高效率 https://zhuanlan.zhihu.com/p/73690883

func TestReadBytes(t *testing.T) {
	file, err := os.Open("writeFile.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	fmt.Println(reader.Size())
	fmt.Println(reader.Buffered())

	// ReadSlice、ReadBytes、ReadString 和 ReadLine 方法
	// 后三个方法最终都是调用ReadSlice来实现的
	// ReadSlice 从输入中读取，直到遇到第一个界定符（delim）为止，返回一个指向缓存中字节的 slice，在下次调用读操作（read）时，这些字节会无效
	data, err := reader.ReadBytes(byte('s'))
	if err != nil {
		panic(err)
	}

	// 返回的结果是包含界定符本身的
	fmt.Println(string(data))
}

// 第一次ReadSlice的结果（line），在第二次调用读操作后，内容发生了变化。
// 也就是说，ReadSlice 返回的 []byte 是指向 Reader 中的 buffer ，而不是 copy 一份返回。
//
// 正因为ReadSlice 返回的数据会被下次的 I/O 操作重写，因此许多的客户端会选择使用 ReadBytes 或者 ReadString 来代替
func TestReadSlice(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))

	line, _ := reader.ReadSlice('\n')    // 把这里和下面换成ReadString看看
	fmt.Printf("the line:%s\n", line)

	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
}

// 如果 ReadSlice 在找到界定符之前遇到了 error ，它就会返回缓存中所有的数据和错误本身（经常是 io.EOF）
// 如果在找到界定符之前缓存已经满了，ReadSlice 会返回 bufio.ErrBufferFull 错误
// 当且仅当返回的结果（line）没有以界定符结束的时候，ReadSlice 返回err != nil
// 也就是说，如果ReadSlice 返回的结果 line 不是以界定符 delim 结尾, 那么返回的 err也一定不等于 nil（可能是bufio.ErrBufferFull或io.EOF）
func TestFullBuf(t *testing.T) {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com"),16)

	line, err := reader.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)

	line, err = reader.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)
}

// ReadLine 尝试返回单独的行，不包括行尾的换行符。如果输入中没有行尾标识符，不会返回任何指示或者错误。
// 如果一行大于缓存，isPrefix 会被设置为 true，同时返回该行的开始部分（等于缓存大小的部分）。
// 该行剩余的部分就会在下次调用的时候返回。当下次调用返回该行剩余部分时，isPrefix 将会是 false 。
//
// 跟 ReadSlice 一样，返回的 line 只是 buffer 的引用，在下次执行IO操作时，line 会无效
//
// 注意，返回值中，要么 line 不是 nil，要么 err 非 nil，两者不会同时非 nil。
//
// 读取一行，通常会选择 ReadBytes 或 ReadString。
// 不过，正常人的思维，应该用 ReadLine，只是不明白为啥 ReadLine 的实现不是通过 ReadBytes，然后清除掉行尾的\n（或\r\n）
// 它现在的实现，用不好会出现意想不到的问题，比如丢数据。
func TestReadLine(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))

	line, isPrefix, err := reader.ReadLine()
	fmt.Printf("line:%s\tisPrefix:%v\terror:%v\n", line, isPrefix, err)

	line, isPrefix, err = reader.ReadLine()
	fmt.Printf("line:%s\tisPrefix:%v\terror:%v\n", line, isPrefix, err)

	line, isPrefix, err = reader.ReadLine()
	fmt.Printf("line:%s\tisPrefix:%v\terror:%v\n", line, isPrefix, err)
}

// 从方法的名称可以猜到，该方法只是“窥探”一下 Reader 中没有读取的 n 个字节。好比栈数据结构中的取栈顶元素，但不出栈。
// Peek返回的 []byte 只是 buffer 中的引用，在下次IO操作后会无效
// 可见该方法（以及ReadSlice这样的，返回buffer引用的方法）对多 goroutine 是不安全的，也就是在多并发环境下，不能依赖其结果
func TestPeek(t *testing.T) {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
	line1, _ := reader.Peek(14)
	fmt.Printf("%s\n", line1)
	line2, _ := reader.Peek(14)
	fmt.Printf("%s\n", line2)
}


func TestWriteBuf(t *testing.T) {
	file, err := os.Create("writer.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	// Available 方法获取缓存中还未使用的字节数（缓存大小 - 字段 n 的值）；Buffered 方法获取写入当前缓存中的字节数（字段 n 的值）
	fmt.Println(writer.Buffered(), writer.Size(), writer.Available())

	n, err := writer.Write([]byte("hello"))
	if err != nil {
		panic(err)
	}

	fmt.Println(n, writer.Buffered(), writer.Size(), writer.Available())
	// 该方法将缓存中的所有数据写入底层的 io.Writer 对象中。
	// 使用 bufio.Writer 时，在所有的 Write 操作完成之后，应该调用 Flush 方法使得缓存都写入 io.Writer 对象中
	writer.Flush()  // 注释掉这行 看看会出现什么问题
}













