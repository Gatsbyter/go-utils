package leecode

import (
	"testing"
)

// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
func Test22(t *testing.T) {

}

// 笨比方法
func largestRectangleArea(heights []int) int {
	length, max := len(heights), 0
	m := make(map[int]bool)

	for i, _ := range heights {
		if m[i] {
			continue
		}

		tmp := 1
		for j := i - 1; j >= 0; j-- {
			if heights[j] < heights[i] {
				break
			}

			if heights[j] == heights[i] {
				m[j] = true
			}

			tmp++
		}

		for j := i + 1; j <= length-1; j++ {
			if heights[j] < heights[i] {
				break
			}

			if heights[j] == heights[i] {
				m[j] = true
			}

			tmp++
		}

		if tmp*heights[i] > max {
			max = tmp * heights[i]
		}
	}

	return max
}

// 单调栈方法
func largestRectangleArea2(heights []int) int {
	return 0
}
