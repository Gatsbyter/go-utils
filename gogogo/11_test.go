package gogogo

import (
	"fmt"
	"testing"
	"time"
)

func Test11(t *testing.T) {

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		for i := 1; ; i+=2 {
			<- ch2
			fmt.Printf("%d", i)
			fmt.Printf("%d", i+1)
			ch1 <- struct{}{}
		}
	}()

	go func() {
		for i := 65; i <= 90; i+=2 {
			<- ch1
			fmt.Printf("%c", rune(i))
			fmt.Printf("%c", rune(i+1))
			ch2 <- struct{}{}
		}
	}()

	time.Sleep(time.Second)

	ch2 <- struct{}{}

	time.Sleep(time.Second)
}
