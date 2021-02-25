//给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

package leecode

import (
	"testing"
)

var matrix = [][]byte{
	{1, 0, 1, 0, 0},
	{1, 0, 1, 1, 1},
	{1, 1, 1, 1, 1},
	{1, 0, 0, 1, 0},
}

func Test_6(t *testing.T) {
	t.Log(maximalRectangle(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	return 0
}
