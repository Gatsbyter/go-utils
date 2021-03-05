package reflect

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// 深度相等 测试
func TestDeepEqual(t *testing.T) {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	fmt.Println(reflect.DeepEqual(got, want)) // true

	// 尽管DeepEqual函数很方便，而且可以支持任意的数据类型，但是它也有不足之处。
	// 例如，它将一个nil值的map和非nil值但是空的map视作不相等，同样nil值的slice 和非nil但是空的slice也视作不相等
	var a, b []string = nil, []string{}
	fmt.Println(reflect.DeepEqual(a, b)) // "false"

	var c, d map[string]int = nil, make(map[string]int)
	fmt.Println(reflect.DeepEqual(c, d)) // "false"
}
