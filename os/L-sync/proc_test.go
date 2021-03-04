package L_sync

import (
	"fmt"
	"runtime"
	"testing"
)

// Go的调度器使用了一个叫做GOMAXPROCS的变量来决定会有多少个操作系统的线程同时执行Go的代码。
// 其默认的值是运行机器上的CPU的核心数，所以在一个有8个核心的机器上时，调度器一次会在8个OS线程上去调度GO代码。
// (GOMAXPROCS是前面说的m:n调度中的n)。
// 在休眠中的或者在通信中被阻塞的goroutine是不需要一个对应的线程来做调度的。
// 在I/O中或系统调用中或调用非Go语言函数时，是需要一个对应的操作系统线程的，但是GOMAXPROCS并不需要将这几种情况计算在内。
//
// 你可以用GOMAXPROCS的环境变量来显式地控制这个参数，
// 或者也可以在运行时用runtime.GOMAXPROCS函数来修改它。
func TestProc(t *testing.T) {
	runtime.GOMAXPROCS(2)

	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}

// 在第一次执行时，最多同时只能有一个goroutine被执行。
// 初始情况下只有main goroutine被执行，所以会打印很多1。
// 过了一段时间后，GO调度器会将其置为休眠，并唤醒另一个goroutine，这时候就开始打印很多0了，
// 在打印的时候，goroutine是被调度到操作系统线程上的。
//
// 在第二次执行时，我们使用了两个操作系统线程，
// 所以两个goroutine可以一起被执行，以同样的频率交替打印0和1。
//
// 我们必须强调的是goroutine的调度是受很多因子影响的，
// 而runtime也是在不断地发展演进的，
// 所以这里的你实际得到的结果可能会因为版本的不同而与我们运行的结果有所不同。
