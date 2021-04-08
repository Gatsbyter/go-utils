package leecode

//https://leetcode-cn.com/problems/minimum-path-sum/
//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j != 0 {
				grid[i][j] += grid[i][j-1]
			}

			if i != 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			}

			if i != 0 && j != 0 {
				grid[i][j] += Min(grid[i-1][j], grid[i][j-1])
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
