package L_sync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var count int32

// 错误示例

func add(wg *sync.WaitGroup) {
	defer wg.Done()
	count++
	// count 其实分三步
	// 1、cpu从内存读count
	// 2、cpu执行count = count + 1
	// 3、cpu将count写回内存
}

func addAtomic(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(&count, 1)
}

func TestAtomic(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		//go add(&wg)
		go addAtomic(&wg)
	}
	wg.Wait()
	fmt.Println(count)
}
