package gogogo

import (
	"fmt"
	"testing"
)

func Test10(t *testing.T) {
	person := &Person{28}

	defer fmt.Println(person.age)

	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}

type Person struct {
	age int
}
