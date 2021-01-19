package main

import (
	"flag"
	"fmt"
)

func main() {
	num := flag.Int("n", 10, "number")
	flag.Parse()
	fmt.Println(*num)
}
