package L_time

import (
	"fmt"
	"testing"
	"time"
)

// 自行体会
func TestTimer(t *testing.T) {
	c := make(chan int)

	go func() {
		// time.Sleep(1 * time.Second)
		time.Sleep(3 * time.Second)
		<-c
	}()

	select {
	case c <- 1:
		fmt.Println("channel...")
	case <-time.After(2 * time.Second):
		close(c)
		fmt.Println("timeout...")
	}
}

func TestTimerFunc(t *testing.T) {
	start := time.Now()
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("after func callback, elaspe:", time.Now().Sub(start))
	})

	time.Sleep(1 * time.Second)
	// time.Sleep(3*time.Second)
	// Reset 在 Timer 还未触发时返回 true；触发了或 Stop 了，返回 false
	if timer.Reset(3 * time.Second) {
		fmt.Println("timer has not trigger!")
	} else {
		fmt.Println("timer had expired or stop!")
	}

	time.Sleep(10 * time.Second)
}
