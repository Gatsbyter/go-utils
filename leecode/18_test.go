// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
// https://leetcode-cn.com/problems/longest-valid-parentheses/

package leecode

import (
	"fmt"
	"testing"
)

func Test18(t *testing.T) {
	fmt.Println(longestValidParentheses("()(())"))
}

func longestValidParentheses(s string) int {
	length := len(s)

	if length < 2 {
		return 0
	}

	max := 0
	dp := make([]int, length+1)
	dp[0], dp[1] = 0, 0

	for i := 2; i < length+1; i++ {
		if s[i-1] == '(' {
			dp[i] = 0
		} else if s[i-2] == '(' {
			dp[i] = dp[i-2] + 2
		} else if i-dp[i-1]-2 >= 0 && s[i-dp[i-1]-2] == '(' {
			dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
		} else {
			dp[i] = 0
		}

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
