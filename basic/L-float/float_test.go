package L_float

import (
	"fmt"
	"math"
	"testing"
)

func TestFloat(t *testing.T) {
	// 下面代码打印e的幂，打印精度是小数点后三个小数精度和8个字符宽度：
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	// math包中除了提供大量常用的数学函数外，还提供了IEEE754浮点数标准中定义的特殊值的创建和测试
	// 正无穷大和负无穷大，分别用于表示太大溢出的数字和除零的结果
	// 还有NaN非数，一般用于表示无效的除法操作结果0/0或Sqrt(-1)
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	// 因为NaN和任何数都是不相等的
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}
