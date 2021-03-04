package benchmark

import (
	"github.com/Gatsbyter/go-utils/tools/hash"
	"testing"
)

// 和普通测试不同的是，默认情况下不运行任何基准测试。
// 我们需要通过-bench命令行标志参数手工指定要运行的基准测试函数
// go test -bench=.  匹配所有基准测试函数

// 基准测试是测量一个程序在固定工作负载下的性能。
// 在Go语言中，基准测试函数和普通测试函数写法类似，但是以Benchmark为前缀名，并且带有一个*testing.B类型的参数；
// *testing.B参数除了提供和*testing.T类似的方法，还有额外一些和性能测量相关的方法。
// 它还提供了一个整数N，用于指定操作执行的循环次数。
func BenchmarkSpecial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hash.Hash("abc")
	}
}

// 运行后的输出是这样的
// BenchmarkSpecial-8   	 7465891	       160.4 ns/op
// -8          : 对应的GOMAXPROCS的值
// 7465891     : 执行了7465891次
// 160.4 ns/op : 平均每次执行花费160.4 ns

// ⚠️ 用这个 go test -bench=. -benchmem 还可以有更多信息
