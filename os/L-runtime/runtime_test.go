package L_runtime

import (
	"fmt"
	"runtime"
	"testing"
)

func TestRuntime(t *testing.T) {

	// runtime.Gosched() // 当前routine让出cpu 但不会挂起
	// runtime.Goexit()  // 退出当前routine
	// runtime.NumGoroutine()  // 返回当前存活的routine
	runtime.GC() // 让系统强制GC一下

	fmt.Println(runtime.NumCPU()) // cpu核数
	fmt.Println(runtime.GOROOT()) // GOROOT
	fmt.Println(runtime.GOOS)     // 当前操作系统
}
