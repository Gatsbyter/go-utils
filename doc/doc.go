package doc

// defer 是FILO的
// defer 调用非常损耗性能

// golang 和 c 语言不一样，栈区分配的存储空间不会随着函数的返回而释放，本地变量地址所占据的存储空间会生存下来
