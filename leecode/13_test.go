// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

package leecode

import (
	"math"
	"testing"
)

func Test13(t *testing.T) {

}

// 这是我自己想的笨比方法
func maxProduct(nums []int) int {
	dp := make([]int, len(nums))

	for i, v := range nums {
		if i == 0 {
			dp[i] = v
			continue
		}

		max, tmp := v, v
		for j := i - 1; j >= 0; j-- {
			tmp *= nums[j]
			if tmp > max {
				max = tmp
			}
		}

		if dp[i-1] > max {
			dp[i] = dp[i-1]
		} else {
			dp[i] = max
		}
	}

	return dp[len(nums)-1]
}

// 牛逼的方法
func maxProduct2(nums []int) int {
	max, imax, imin := math.MinInt64, 1, 1

	for _, v := range nums {
		if v < 0 {
			imax, imin = imin, imax
		}

		imax = Max(v, imax*v)
		imin = Min(v, imin*v)
		max = Max(max, imax)
	}

	return max
}
