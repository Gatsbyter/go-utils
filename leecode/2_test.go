// 动态规划

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

package leecode

import (
	"math"
	"testing"
)

var nums = []int{10, 9, 2, 5, 3, 7, 101, 18}

func Test_2(t *testing.T) {
	t.Log(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	length := len(nums)
	dp := make([]int, length)

	for i := 0; i < length; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return maxArr(dp)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxArr(arr []int) int {
	max := math.MinInt32
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}
