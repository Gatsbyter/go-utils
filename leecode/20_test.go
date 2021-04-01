// 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数

package leecode

import "math"

func minDistance(word1 string, word2 string) int {
	length1 := len(word1)
	length2 := len(word2)

	dp := make([][]int, length1+1)
	for i, _ := range dp {
		dp[i] = make([]int, length2+1)
	}

	for i := 0; i < length1+1; i++ {
		dp[i][0] = i
	}

	for i := 0; i < length2+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < length1+1; i++ {
		for j := 1; j < length2+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[length1][length2]
}

func min(nums ...int) int {
	min := math.MaxInt64

	for _, i := range nums {
		if min > i {
			min = i
		}
	}

	return min
}
