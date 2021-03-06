package L_sync

// ⚠️ 线程和routine的区别

// 每一个OS线程都有一个固定大小的内存块(一般会是2MB)来做栈，
// 这个栈会用来存储当前正在被调用或挂起(指在调用其它函数时)的函数的内部变量。
//
// 这个固定大小的栈同时很大又很小。
// 因为2MB的栈对于一个小小的goroutine来说是很大的内存浪费，
// 比如对于我们用到的，一个只是用来WaitGroup之后关闭channel的goroutine来说。
// 而对于go程序来说，同时创建成百上千个goroutine是非常普遍的，
// 如果每一个goroutine都需要这么大的栈的话，那这么多的goroutine就不太可能了。
//
// 除去大小的问题之外，固定大小的栈对于更复杂或者更深层次的递归函数调用来说显然是不够的。
// 修改固定的大小可以提升空间的利用率允许创建更多的线程，并且可以允许更深的递归调用，不过这两者是没法同时兼备的。
//
// ------------------------------------------------------------------------------------------------------------
//
// 相反，一个goroutine会以一个很小的栈开始其生命周期，一般只需要2KB。
// 一个goroutine的栈，和操作系统线程一样，会保存其活跃或挂起的函数调用的本地变量，
// 但是和OS线程不太一样的是一个goroutine的栈大小并不是固定的；
// 栈的大小会根据需要动态地伸缩。
//
// 而goroutine的栈的最大值有1GB，比传统的固定大小的线程栈要大得多，
// 尽管一般情况下，大多goroutine都不需要这么大的栈。

// ------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------

// OS线程会被操作系统内核调度。
// 每几毫秒，一个硬件计时器会中断处理器，这会调用一个叫作scheduler的内核函数。
// 这个函数会挂起当前执行的线程并保存内存中它的寄存器内容，检查线程列表并决定下一次哪个线程可以被运行，
// 并从内存中恢复该线程的寄存器信息，然后恢复执行该线程的现场并开始执行线程。
//
// 因为操作系统线程是被内核所调度，所以从一个线程向另一个“移动”需要完整的上下文切换，
// 也就是说，保存一个用户线程的状态到内存，恢复另一个线程的到寄存器，然后更新调度器的数据结构。
// 这几步操作很慢，因为其局部性很差需要几次内存访问，并且会增加运行的cpu周期。
//
// ------------------------------------------------------------------------------------------------------------
//
// Go的运行时包含了其自己的调度器，这个调度器使用了一些技术手段，比如m:n调度，
// 因为其会在n个操作系统线程上多工(调度)m个goroutine。
// Go调度器的工作和内核的调度是相似的，但是这个调度器只关注单独的Go程序中的goroutine(译注：按程序独立)。
//
// 和操作系统的线程调度不同的是，
// Go调度器并不是用一个硬件定时器而是被Go语言"建筑"本身进行调度的。
// 例如当一个goroutine调用了time.Sleep或者被channel调用或者mutex操作阻塞时，
// 调度器会使其进入休眠并开始执行另一个goroutine直到时机到了再去唤醒第一个goroutine。
// 因为这种调度方式不需要进入内核的上下文，所以重新调度一个goroutine比调度一个线程代价要低得多。
