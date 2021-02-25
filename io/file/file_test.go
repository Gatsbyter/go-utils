package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	file, err := os.Create("test.log")
	if err != nil {
		panic(err)
	}

	n, err := file.Write([]byte("nimabi"))
	if err != nil {
		panic(err)
	}

	t.Log(n)

	err = file.Close()
	if err != nil {
		panic(err)
	}
}

func TestAppend(t *testing.T) {
	// 第二个参数是打开时的属性
	// 第三个参数是权限模式
	file, err := os.OpenFile("test.log", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	n, err := file.Write([]byte("\nfuck"))
	if err != nil {
		panic(err)
	}

	t.Log(n)

	err = file.Close()
	if err != nil {
		panic(err)
	}
}

func TestTruncate(t *testing.T) {
	err := os.Truncate("test.log", 7)
	if err != nil {
		panic(err)
	}
}

func TestFileStat(t *testing.T) {
	fileInfo, err := os.Stat("test.log")
	if err != nil {
		panic(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

func TestCopy(t *testing.T) {
	originalFile, err := os.Open("test.log")
	if err != nil {
		panic(err)
	}
	defer originalFile.Close()

	newFile, err := os.Create("test_copy.log")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		panic(err)
	}

	t.Logf("Copied %d bytes.", bytesWritten)
}

func TestSeek(t *testing.T) {
	file, err := os.Open("test.log")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// 偏离位置，可以是正数也可以是负数
	var offset int64 = 5
	// whence用来计算offset的初始位置
	// 0 = 文件开始位置
	// 1 = 当前位置
	// 2 = 文件结尾处
	whence := 0

	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		panic(err)
	}

	t.Logf("moved to %d\n", newPosition)

	// 从当前位置回退两个字节
	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		panic(err)
	}

	t.Logf("moved to %d\n", newPosition)

	// 寻找当前位置
	currentPosition, err := file.Seek(0, 1)
	if err != nil {
		panic(err)
	}

	t.Logf("Current position %d\n", currentPosition)

	// 转到文件开始处
	newPosition, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	t.Logf("Position after seeking %d\n", newPosition)
}

func TestWrite(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	//currentPosition, err := file.Seek(0, 2)
	//if err != nil {
	//	panic(err)
	//}
	//
	//t.Logf("Current position %d\n", currentPosition)

	bytesWritten, err := file.Write([]byte("Bytes!\n"))
	if err != nil {
		panic(err)
	}

	t.Logf("Wrote %d bytes\n", bytesWritten)
}

func TestFastWrite(t *testing.T) {
	err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		panic(err)
	}
}
