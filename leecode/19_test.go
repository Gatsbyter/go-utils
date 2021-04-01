//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

package leecode

func maxSubArray(nums []int) int {
	max, last := 0, 0

	for i, v := range nums {
		if i == 0 {
			max, last = v, v
			continue
		}

		if last <= 0 {
			last = v
		} else {
			last += v
		}

		if last > max {
			max = last
		}
	}

	return max
}
