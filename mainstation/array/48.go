package array

func rotate(matrix [][]int) {
	/*
		n*n矩阵的旋转可以拆分为两步：1.中轴对称翻转 2.主对角线翻转
		顺时针转90度：1.以横中轴为轴翻转 2.以主对角线为轴翻转
		逆时针转90度：1.以纵中轴为轴翻转 2.以主对角线为轴翻转
	*/

	// 1.矩阵的维度n
	n := len(matrix)

	// 2.先以横中轴为轴翻转
	top := 0
	button := n - 1
	for top < button {
		for i := 0; i < n; i++ {
			matrix[top][i], matrix[button][i] = matrix[button][i], matrix[top][i]
		}
		top++
		button--
	}

	// 3.再以主对角线为轴翻转
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
