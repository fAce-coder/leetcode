package array

func maxValue(grid [][]int) int {
	// 动态规划

	/*
	   建立一个动态规划数组
	   数组中的每个位置储存起点到这个位置的最大礼物价值
	   除了第一行和第一列，每一个位置的值，都是max(上面来的数，左面来的数)+当前数
	   此时最右下角的数就是礼物的最大价值
	*/

	// 建立一个动态规划数组,数组的大小和原数组相同
	var raw int = len(grid)
	var col int = len(grid[0])
	var matrix [][]int = make([][]int, raw)
	for i := 0; i < raw; i++ {
		matrix[i] = make([]int, col)
	}
	// 遍历动态数组，对其进行动态规划
	for i := 0; i < raw; i++ {
		for j := 0; j < col; j++ {
			// 当第一行第一列，即起始位置，则礼物价值就是原始数组该位置的值
			if i == 0 && j == 0 {
				matrix[i][j] = grid[i][j]
				// 当第一行，只能从左边来
			} else if i == 0 && j != 0 {
				matrix[i][j] = matrix[i][j-1] + grid[i][j]
				// 当第一列，只能从上边来
			} else if i != 0 && j == 0 {
				matrix[i][j] = matrix[i-1][j] + grid[i][j]
				// 其他情况，max(上面来的数，左面来的数)+当前数
			} else {
				if matrix[i][j-1] >= matrix[i-1][j] {
					matrix[i][j] = matrix[i][j-1] + grid[i][j]
				} else {
					matrix[i][j] = matrix[i-1][j] + grid[i][j]
				}
			}
		}
	}
	// 此时最右下角的数就是礼物的最大价值
	return matrix[raw-1][col-1]
}
