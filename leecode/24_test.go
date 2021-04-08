package leecode

//https://leetcode-cn.com/problems/word-break/submissions/
//给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	dp := make([]bool, len(s)+1)

	for _, v := range wordDict {
		m[v] = true
	}

	dp[0] = true

	for i := range s {
		for j := i; j >= 0; j-- {
			if dp[j] && m[s[j:i+1]] {
				dp[i+1] = true
				break
			}
		}
	}

	return dp[len(s)]
}
