package L_bytes

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	// bytes 有很多便捷方法 需要的时候用就行了
}

// bytes.Buffer 实现了 io 包下的 ByteScanner, ByteWriter, ReadWriter, Reader, ReaderFrom,
// RuneReader, RuneScanner, StringWriter, Writer, WriterTo 等接口，可以方便的进行读写操作。
//
// 对象可读取数据为 buf[off : len(buf)], off 表示进度下标，lastRead 表示最后读取的一个字符所占字节数，方便 Unread* 相关操作
func TestBuffer(t *testing.T) {
	// 初始化
	a := bytes.NewBufferString("Hello World")
	b := bytes.NewBuffer([]byte("Hello World"))
	c := bytes.Buffer{}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 没啥好说的 当成io淦就是了
}
