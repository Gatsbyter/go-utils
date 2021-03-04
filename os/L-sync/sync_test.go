package L_sync

import (
	"testing"
)

// Go的口头禅“不要使用共享数据来通信；使用通信来共享数据”

// 数据竞争会在两个以上的goroutine并发访问相同的变量 且至少其中一个为写操作时发生。
// 避免数据竞争：

func TestSync(t *testing.T) {

}

// 只要在go build，go run或者go test命令后面加上-race的flag，
// 就会使编译器创建一个你的应用的“修改”版或者一个附带了能够记录所有运行期对共享变量访问工具的test，
// 并且会记录下每一个读或者写共享变量的goroutine的身份信息。
//
// 另外，修改版的程序会记录下所有的同步事件，比如go语句，channel操作，
// 以及对(*sync.Mutex).Lock，(*sync.WaitGroup).Wait等等的调用。
