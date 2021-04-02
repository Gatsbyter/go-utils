package leecode

// https://leetcode-cn.com/problems/maximal-rectangle
func maximalRectangle(matrix [][]byte) int {
	max := 0

	help := make([][]int, len(matrix)) // 用来记录每个点左边最多多少个1
	for i, _ := range help {
		help[i] = make([]int, len(matrix[i]))
	}

	for i, row := range matrix {
		var tmp = 0
		for j, _ := range row {
			if matrix[i][j] == '1' {
				tmp++
			} else {
				tmp = 0
			}

			help[i][j] = tmp
			if tmp > max {
				max = tmp
			}

			width := help[i][j]
			for k := i - 1; k >= 0 && help[k][j] > 0; k-- {
				if help[k][j] < width {
					width = help[k][j]
				}

				area := width * (i - k + 1)
				if area > max {
					max = area
				}
			}
		}
	}

	return max
}
