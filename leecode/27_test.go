package leecode

func numIslands(grid [][]byte) int {
	width := len(grid)
	length := len(grid[0])

	fa := make([][]int, width)
	for i := range fa {
		fa[i] = make([]int, length)
	}

	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {

		}
	}

	return 0
}
