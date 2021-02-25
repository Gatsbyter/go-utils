//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水

package leecode

import (
	"testing"
)

var height = []int{4, 2, 0, 3, 2, 5}

func Test_5(t *testing.T) {
	t.Log(trap(height))
}

func trap(height []int) int {
	length := len(height)

	if length == 0 {
		return 0
	}

	maxHeight := 0
	maxIndex := 0
	for i, h := range height {
		if h > maxHeight {
			maxHeight = h
			maxIndex = i
		}
	}

	res := 0
	lastMax := height[0]

	for i := 1; i < maxIndex; i++ {
		if height[i] < lastMax {
			res += lastMax - height[i]
			continue
		}

		lastMax = height[i]
	}

	lastMax = height[length-1]

	for i := length - 1; i > maxIndex; i-- {
		if height[i] < lastMax {
			res += lastMax - height[i]
			continue
		}

		lastMax = height[i]
	}

	return res
}
