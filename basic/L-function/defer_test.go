package L_function

import (
	"fmt"
	"log"
	"testing"
	"time"
)

// 调试复杂程序时，defer机制也常被用于记录何时进入和退出函数
// 这个还行 一般牛逼
func TestDefer(t *testing.T) {
	defer trace("TestDefer")() // 这个的括号很重要
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

// 被延迟执行的匿名函数甚至可以修改函数返回给调用者的返回值
func TestTriple(t *testing.T) {
	fmt.Println(triple(4)) // "12"
}

func double(x int) int {
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
