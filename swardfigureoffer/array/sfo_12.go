package array

func exist(board [][]byte, word string) bool {
	// 递归 + 回溯

	// 左右边界
	var raw int = len(board) - 1
	var col int = len(board[0]) - 1

	// 递归函数
	var inner func(i int, j int, k int) bool
	inner = func(i int, j int, k int) bool {
		if (i < 0 || i > raw) || (j < 0 || j > col) || board[i][j] != byte(word[k]) {
			return false
		}
		// 如果当前是最后一个单词，则返回true
		if k == len(word)-1 {
			return true
		}
		// 递归寻找下一个字母
		// 将当前已经找过的字母变成空，防止重复寻找
		temp := board[i][j]
		board[i][j] = ' '
		// 向上下左右四个方向，递归寻找下一个单词
		res := inner(i+1, j, k+1) || inner(i-1, j, k+1) || inner(i, j-1, k+1) || inner(i, j+1, k+1)
		// 回溯，将置为空的字母复原，供其他路径查找
		board[i][j] = temp
		return res
	}

	// 遍历数组，以每个点为起点，尝试找到这个单词
	for i := 0; i <= raw; i++ {
		for j := 0; j <= col; j++ {
			if inner(i, j, 0) {
				return true
			}
		}
	}
	// 如果跳出循环还没找到，说明不存在，返回false
	return false
}
