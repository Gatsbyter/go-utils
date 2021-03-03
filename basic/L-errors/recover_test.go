package L_errors

import (
	"runtime/debug"
	"testing"
)

// ⚠️defer 是FILO的
// defer 调用非常损耗性能

// panic和recover
func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()

	panic("bye I'm dead")
	t.Log("exit")
}

// panic后，函数会一层层的退出，直到出现recover捕获，没有的话就gg
func TestFuncRecover(t *testing.T) {
	defer func() {
		t.Log(recover())
		// 打印错误堆栈
		debug.PrintStack()
	}()

	FuncPanic(t)
}

func FuncPanic(t *testing.T) {
	defer t.Log("11111")
	defer t.Log("22222")

	panic("bye I'm dead")
}

// 连续调用panic，只有最后一个会被捕获
func TestMultiPanic(t *testing.T) {
	defer func() {
		for {
			if err := recover(); err != nil {
				t.Log(err)
			} else {
				t.Log("fuck you")
				return
			}
		}
	}()

	defer func() {
		panic("bye I'm dead")
	}()

	panic("will I be caught")
}
