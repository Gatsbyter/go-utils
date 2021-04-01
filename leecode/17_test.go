//找出这个数组排序出的所有数中，刚好比当前数大的那个数
//
//比如当前 nums = [1,2,3]。这个数是123，找出1，2，3这3个数字排序可能的所有数，排序后，比123大的那个数 也就是132
//
//如果当前 nums = [3,2,1]。这就是1，2，3所有排序中最大的那个数，那么就返回1，2，3排序后所有数中最小的那个，也就是1，2，3 -> [1,2,3]
// https://leetcode-cn.com/problems/next-permutation/
package leecode

func nextPermutation(nums []int) {
	length := len(nums)

	for i := length - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			for j := length - 1; j >= i; j-- {
				if nums[i-1] < nums[j] {
					nums[i-1], nums[j] = nums[j], nums[i-1]

					for p, q := i, length-1; p < q; p, q = p+1, q-1 {
						nums[p], nums[q] = nums[q], nums[p]
					}
					return
				}
			}
		}
	}

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
