package L_sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Once 保证只运行一次

var once sync.Once

func TestOnce(t *testing.T) {
	for i := 0; i < 3; i++ {
		go DoOnce()
	}

	time.Sleep(time.Second * 2)
}

func DoOnce() {
	once.Do(func() {
		fmt.Println("我只能出现一次")
	})
	fmt.Println("我会出现三次")
}
