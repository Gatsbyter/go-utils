package L_time

import (
	"fmt"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	d1 := time.Duration(654321234567890)
	fmt.Println(d1)

	d2, err := time.ParseDuration("2h30m55s")
	if err != nil {
		panic(err)
	}

	fmt.Println("hours    ------- ", d2.Hours())
	fmt.Println("minutes  ------- ", d2.Minutes())
	fmt.Println("seconds  ------- ", d2.Seconds())
	fmt.Println("milli    ------- ", d2.Milliseconds())
}
