//动态规划

//给你两个数组nums1和nums2。
//请你返回 nums1 和 nums2 中两个长度相同的 非空 子序列的最大点积。
//数组的非空子序列是通过删除原数组中某些元素（可能一个也不删除）后剩余数字组成的序列，但不能改变数字间相对顺序。比方说，[2,3,5] 是 [1,2,3,4,5] 的一个子序列而 [1,5,3] 不是。

package leecode

import (
	"math"
	"testing"
)

// == 18
//var nums1 = []int{2,1,-2,5}
//var nums2 = []int{3,0,-6}

// == 21
//var nums1 = []int{3,-2}
//var nums2 = []int{2,-6,7}

// == -1
var nums1 = []int{5, -4, -3}
var nums2 = []int{-4, -3, 0, -4, 2}

func Test_3(t *testing.T) {
	t.Log(maxDotProduct(nums1, nums2))
}

func maxDotProduct(nums1 []int, nums2 []int) int {
	length1 := len(nums1)
	length2 := len(nums2)

	res := make([][]int, length1)
	for i := range res {
		res[i] = make([]int, length2)
	}

	res[0][0] = nums1[0] * nums2[0]

	for i := 1; i < length1; i++ {
		res[i][0] = max(res[i-1][0], nums1[i]*nums2[0])
	}

	for j := 1; j < length2; j++ {
		res[0][j] = max(res[0][j-1], nums1[0]*nums2[j])
	}

	for i := 1; i < length1; i++ {
		for j := 1; j < length2; j++ {
			temp := nums1[i] * nums2[j]
			res[i][j] = max2(res[i-1][j], res[i][j-1], res[i-1][j-1]+temp, temp)
		}
	}

	return maxArr2(res)
}

func max2(nums ...int) int {
	max := nums[0]

	for _, i := range nums {
		if max < i {
			max = i
		}
	}

	return max
}

func maxArr2(arr [][]int) int {
	max := math.MinInt32
	for _, a1 := range arr {
		for _, a2 := range a1 {
			if a2 > max {
				max = a2
			}
		}
	}
	return max
}
