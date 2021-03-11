package doc

// golang 和 c 语言不一样，栈区分配的存储空间不会随着函数的返回而释放，本地变量地址所占据的存储空间会生存下来

// v, ok = m[key]
// v, ok = x.(T)
// v, ok = <- ch

// 在Go语言中，%取模运算符的符号和被取模数的符号总是一致的，因此-5%3和-5%-3结果都是-2。
// 除法运算符/的行为则依赖于操作数是否为全为整数，比如5.0/4.0的结果是1.25，但是5/4的结果是1，因为整数除法会向着0方向截断余数

// 所有以_test为后缀包名的测试外部扩展包都由go test命令独立编译，普通包和测试的外部扩展包是相互独立的

// 包的匿名导入 用 _ 可以不引用 而且会初始化这个包
// 包的 . 导入 直接使用 不要别名

// 在*_test.go文件中，有三种类型的函数：测试函数、基准测试(benchmark)函数、示例函数。
//
// 一个测试函数是以Test为函数名前缀的函数，用于测试程序的一些逻辑行为是否正确；
// go test命令会调用这些测试函数并报告测试结果是PASS或FAIL。
//
// 基准测试函数是以Benchmark为函数名前缀的函数，它们用于衡量一些函数的性能；
// go test命令会多次运行基准函数以计算一个平均的执行时间。
//
// 示例函数是以Example为函数名前缀的函数，提供一个由编译器保证正确性的示例文档。
//
// go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，生成一个临时的main包用于调用相应的测试函数，
// 接着构建并运行、报告测试结果，最后清理测试中生成的临时文件。

// ⚠️GoConvey 是一个很好的测试框架

// 内存分配简述
// 1、每次从操作系统申请一大块内存 以减少系统调用
// 2、将申请到的大内存 按特定大小 预先切割成小块 构造链表
// 3、为对象分配内存时 从大小合适的链表提取一小块
// 4、回收对象内存时 将该小块内存归还给链表
// 5、如果闲置内存较多 归还一部分给操作系统 降低开销

// go的内存分配 直接采用了tcmalloc

// 微对象: 0 - 16B
// 小对象: 16B - 32KB
// 大对象: 32KB - 无穷