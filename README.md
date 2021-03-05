# go-utils

### basic(GoLang基本特性)

- 位操作 [bit](./basic/L-bit)
- 变量(定义、初始化、格式化输出) [var](./basic/L-var)
- 变量生命周期 [life](./basic/L-life)
- 浮点数探索 [float](./basic/L-float)
- 字符串基础(string和rune) [string](./basic/L-string)
- iota基础 [iota](./basic/L-iota)
- 数组&切片 [slice](./basic/L-slice)
- map&结构体 [map](./basic/L-map)
- JSON [json](./basic/L-json)
- 函数&方法(闭包、延迟、传参) [func](./basic/L-function)
- 错误处理 [errors](./basic/L-errors)

### command(获取命令行参数)

- 通过flag库获取命令行参数以及相关操作 [flag](./command/flag/)
- go原生获取命令行参数 [os](./command/os)

### data-struct(数据结构实现)

- 堆(原生) [heap](./data_struct/heap)
- 环(原生) [ring](./data_struct/ring)
- 链表(原生) [list](./data_struct/list)
- 链表 [link_list](./data_struct/link_list)
- 队列 [queue](./data_struct/link_list)
- 栈 [stack](./data_struct/stack)
- 树 [tree](./data_struct/tree)
- 排序 [sort](./data_struct/sort)

### doc(文档)

- 一些笔记
- GO程序设计语言
- 人月神话
- 深入理解计算机系统
- 重构

### leecode(算法题)

### net(网络相关)

- tcp粘包 [sticky_pkg](./net/sticky_pkg)
- tcp发包 [tcp](./net/tcp)
- http相关 [http](./net/http)

### os(操作系统相关)

- 目录&文件相关 [dir](./os/L-dir)
- 环境变量相关 [env](./os/L-env)
- 执行程序 [exec](./os/L-exec)
- 基于共享变量的并发 [sync](./os/L-sync)

### reflect(反射)

- 基本说明 [basic](./reflect/reflect_test.go)
- 类型 [typeOf](./reflect/type_test.go)
- 值 [valueOf](./reflect/value_test.go)
- 获取标签值 [tag](./reflect/tag_test.go)
- 深度相等 [equal](./reflect/equal_test.go)

### routine(协程)

- 缓冲区 [buf](./routine/buf_test.go)
- 通道 [channel](./routine/chan_test.go)
- 通道关闭测试 [close](./routine/close_test.go)
- select测试 [select](./routine/select_test.go)
- routine泄漏 [leak](./routine/leak_test.go)

### service(app框架-处理信号)

### tools(工具)

- 检测通道是否关闭 [chan-close](./tools/chan-close)
- 一键哈希 [hash](./tools/hash)
- 获取本机公网IP [self-ip](./tools/self-ip)
- 类型转换工具 [type_convert](./tools/type_convert)
- 获取当前时间 [tm](./tools/tm)
- 基准测试 [benchmark](./tools/benchmark)

### unsafe(底层玩法)

- 获取内存大小 [sizeof](./unsafe/unsafe_test.go)
- 结构体偏移对齐 [struct](./unsafe/struct_test.go)
- void指针 [pointer](./unsafe/point_test.go)
- 内存打印(查看字节序) [byte](./unsafe/byte_test.go)