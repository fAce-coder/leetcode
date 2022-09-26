package array

func spiralOrder(matrix [][]int) []int {
	// 顺时针打印矩阵

	var res []int = make([]int, 0)
	// 特殊情况
	if len(matrix) == 0 {
		return res
	}
	// 四个边界
	var top int = 0
	var button int = len(matrix) - 1
	var left int = 0
	var right int = len(matrix[0]) - 1

	for {
		// 从左到右
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		if top > button {
			break
		}
		// 从上到下
		for i := top; i <= button; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		// 从右到左
		for i := right; i >= left; i-- {
			res = append(res, matrix[button][i])
		}
		button--
		if button < top {
			break
		}
		// 从下到上
		for i := button; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return res

}
