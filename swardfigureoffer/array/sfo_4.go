package array

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 二维数组查找

	// 因为矩阵从左到右递增，从上到下递增，因此选择左下角或者右上角作为起点
	// 1.左下角作为起点：比它小的数都在它上边，比它大的数都在它右边
	// 2.右上角作为起点：比它小的数都在它左边，比它大的数都在它下边

	// 特殊情况
	if len(matrix) == 0 {
		return false
	}
	// 这里选择左下角作为起点
	var i int = len(matrix) - 1
	var j int = 0
	var right int = len(matrix[0]) - 1
	// 移动索引，寻找目标数
	for i >= 0 && j <= right {
		// 如果目标数大于当前数，就往右边找
		if target > matrix[i][j] {
			j++
			// 如果目标数小于当前数，就往上边找
		} else if target < matrix[i][j] {
			i--
			// 如果目标数等于当前数，就返回true
		} else {
			return true
		}
	}
	// 如果越界还没有找到，说明目标数不在矩阵中，返回false
	return false
}
