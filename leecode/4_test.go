// 图片反转

package leecode

import (
	"testing"
)

var pic [][]int

func Test_4(t *testing.T) {
	t.Log(flipAndInvertImage(pic))
}

func flipAndInvertImage(A [][]int) [][]int {
	res := make([][]int, len(A))

	for i, row := range A {
		length := len(row)
		res[i] = make([]int, length)

		for j := range row {
			res[i][j] = row[length-j]
		}
	}

	return res
}
