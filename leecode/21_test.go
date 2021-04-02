package leecode

// 笨比阶乘
func clumsy(N int) int {
	if N == 1 || N == 2 {
		return 1
	}

	if N == 3 {
		return 6
	}

	res := N * (N - 1) / (N - 2)
	flag, i := true, N-3 // true 是 + false 是 -

	for {
		if flag { // 加法
			res += i
			if i == 1 {
				return res
			}
			i--
		} else { // 剑法
			if i == 1 || i == 2 {
				return res - i
			} else if i == 3 {
				return res - 6
			} else {
				res -= i * (i - 1) / (i - 2)
				i -= 3
			}
		}

		flag = !flag
	}
}
