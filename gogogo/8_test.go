package gogogo

import (
	"fmt"
	"testing"
)

// 当defer被声明时，其参数就会被实时解析
func Test8(t *testing.T) {
	i := 5
	defer fuck(i)
	i = i + 10
}

func fuck(i int) {
	fmt.Println(i)
}
