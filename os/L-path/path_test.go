package L_path

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// ⚠️ !!!!!!!!!!!!!!!!! 兼容操作系统的文件路径操作

// path/filepath 包涉及到路径操作时，路径分隔符使用 os.PathSeparator。
// 不同系统，路径表示方式有所不同，比如 Unix 和 Windows 差别很大。
// 本包能够处理所有的文件路径，不管是什么系统。
//
// 注意，路径操作函数并不会校验路径是否真实存在。

func TestPath(t *testing.T) {
	path1 := "/Users/my/fuck/nope.go"
	path2 := "abc"
	path3 := ""

	fmt.Println(filepath.Base(path1)) // 返回文件
	fmt.Println(filepath.Dir(path1))  // 返回路径

	fmt.Println(filepath.Base(path2)) // abc
	fmt.Println(filepath.Dir(path2))  // .

	fmt.Println(filepath.Base(path3)) // .
	fmt.Println(filepath.Dir(path3))  // .
}

func TestAbs(t *testing.T) {
	// Abs底层调用的是os.Getwd() 然后拼装
	// 在 os.Getwd 出错时，Abs 会返回该错误，一般不会出错，如果路径名长度超过系统限制，则会报错
	fmt.Println(filepath.Abs("./abc"))

	abs, _ := os.Getwd()
	fmt.Println(abs)
	fmt.Println(filepath.IsAbs(abs))
}

// Rel 函数返回一个相对路径，将 basepath 和该路径用路径分隔符连起来的新路径在词法上等价于 targpath。
// 也就是说，Join(basepath, Rel(basepath, targpath)) 等价于 targpath
func TestRel(t *testing.T) {
	fmt.Println(filepath.Rel("/home/me/go", "/data/hello"))

	abs, _ := os.Getwd()
	fmt.Println(filepath.Rel("/Users/rayholu", abs))
}

// Split 函数根据最后一个路径分隔符将路径 path 分隔为目录和文件名两部分（dir 和 file）
// 如果路径中没有路径分隔符，函数返回值 dir 为空字符串，file 等于 path；
// 反之，如果路径中最后一个字符是 /，则 dir 等于 path，file 为空字符串。
// 返回值满足 path == dir+file。dir 非空时，最后一个字符总是 /
func TestSplit(t *testing.T) {
	fmt.Println(filepath.Split("/home/me/go"))
	fmt.Println(filepath.Split("/home/me/go/"))
	fmt.Println(filepath.Split("go"))
}

// Join 函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加路径分隔符。
// 结果是经过 Clean 的，所有的空字符串元素会被忽略。
// 对于拼接路径的需求，我们应该总是使用 Join 函数来处理。
func TestJoin(t *testing.T) {
	fmt.Println(filepath.Join("/home/me/go/", "go", "/home/me/go/"))
}

// Clean 函数通过单纯的词法操作返回和 path 代表同一地址的最短路径。
//
// 1. 将连续的多个路径分隔符替换为单个路径分隔符
// 2. 剔除每一个 . 路径名元素（代表当前目录）
// 3. 剔除每一个路径内的 .. 路径名元素（代表父目录）和它前面的非 .. 路径名元素
// 4. 剔除开始于根路径的 .. 路径名元素，即将路径开始处的 /.. 替换为 /（假设路径分隔符是 /）
func TestClean(t *testing.T) {
	fmt.Println(filepath.Clean("/.././go/jj/../././kk"))
}

// 很牛逼 把指定目录下的路径都打出来
func TestWalk(t *testing.T) {
	err := filepath.Walk("/Users/rayholu/learning/go-utils", walkFunc)
	if err != nil {
		t.Error(err)
	}
}

func walkFunc(path string, info fs.FileInfo, err error) error {
	fmt.Println(path)
	return err
}
